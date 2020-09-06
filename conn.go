package go_mc_wrapped

import (
	"bytes"
	"errors"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
	"io"
	"math"
	"reflect"
	"strconv"
)

type Conn struct {
	net.Conn
}

func Wrap(conn net.Conn) Conn {
	return Conn{conn}
}

var (
	UUID     = reflect.TypeOf(uuid.UUID{}).Kind()
	POSITION = reflect.TypeOf(pk.Position{}).Kind()
)

func writeVarInt(pak *pk.Packet, num uint32) {
	for {
		b := num & 0x7F
		num >>= 7
		if num != 0 {
			b |= 0x80
		}
		pak.Data = append(pak.Data, byte(b))
		if num == 0 {
			break
		}
	}
}

func writeInt(pak *pk.Packet, n uint32) {
	pak.Data = append(pak.Data, byte(n>>24), byte(n>>16), byte(n>>8), byte(n))
}

func writeLong(pak *pk.Packet, n uint64) {
	pak.Data = append(pak.Data, byte(n>>56), byte(n>>48), byte(n>>40), byte(n>>32),
		byte(n>>24), byte(n>>16), byte(n>>8), byte(n))
}

func (that *Conn) Send(packet interface{}) error {
	type1, value := reflect.TypeOf(packet), reflect.ValueOf(packet)
	pak := pk.Packet{}
	field1 := type1.Field(0)
	val, err := strconv.ParseInt(field1.Tag.Get("id"), 10, 8)
	if err != nil {
		return err
	}
	pak.ID = byte(val)
	if field1.Name != "PacketId" {
		for i, num := 0, type1.NumField(); i < num; i++ {
			f, kind := value.Field(i), type1.Field(i).Type
			switch kind.Kind() {
			case reflect.Bool:
				if f.Bool() {
					pak.Data = append(pak.Data, 1)
				} else {
					pak.Data = append(pak.Data, 0)
				}
			case reflect.String:
				byteString := []byte(f.String())
				writeVarInt(&pak, uint32(len(byteString)))
				pak.Data = append(pak.Data, byteString...)
			case reflect.Uint8:
				pak.Data = append(pak.Data, byte(f.Uint()))
			case reflect.Int8:
				pak.Data = append(pak.Data, byte(f.Int()))
			case reflect.Int16:
				n := uint16(f.Int())
				pak.Data = append(pak.Data, byte(n>>8), byte(n))
			case reflect.Uint16:
				t := f.Uint()
				pak.Data = append(pak.Data, byte(t>>8), byte(t))
			case reflect.Int:
				writeVarInt(&pak, uint32(f.Int()))
			case reflect.Int32:
				writeInt(&pak, uint32(f.Int()))
			case reflect.Int64:
				n := uint64(f.Int())
				if val, ok := type1.Field(i).Tag.Lookup("var"); ok && val == "true" {
					for {
						b := n & 0x7F
						n >>= 7
						if n != 0 {
							b |= 0x80
						}
						pak.Data = append(pak.Data, byte(b))
						if n == 0 {
							break
						}
					}
				} else {
					writeLong(&pak, n)
				}
			case reflect.Float32:
				writeInt(&pak, math.Float32bits(float32(f.Float())))
			case reflect.Float64:
				writeLong(&pak, math.Float64bits(f.Float()))
			case UUID:
				if f.IsNil() {
					break
				}
				t := f.Interface().(uuid.UUID)
				pak.Data = append(pak.Data, t[0], t[1], t[2], t[3], t[4], t[5], t[6], t[7], t[8], t[9], t[10], t[11],
					t[11], t[12], t[13], t[14], t[15])
			case POSITION:
				if f.IsNil() {
					break
				}
				t := f.Interface().(pk.Position)
				position := uint64(t.X&0x3FFFFFF)<<38 | uint64((t.Z&0x3FFFFFF)<<12) | uint64(t.Y&0xFFF)
				writeLong(&pak, position)
			case reflect.Array:
				if kind.Elem().Kind() == reflect.Uint8 {
					pak.Data = append(pak.Data, f.Bytes()...)
				}
			}
		}
	}
	return that.WritePacket(pak)
}

func readVarInt(r *bytes.Reader) (int, error) {
	var n uint32
	for i := 0; ; i++ {
		sec, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		n |= uint32(sec&0x7F) << uint32(7*i)
		if i >= 5 {
			return 0, errors.New("VarInt is too big")
		} else if sec&0x80 == 0 {
			break
		}
	}
	return int(n), nil
}

func readInt(r *bytes.Reader) (int32, error) {
	data := make([]byte, 4)
	index, err := r.Read(data)
	if err != nil || index != 4 {
		return 0, err
	}
	return int32(data[0])<<24 | int32(data[1])<<16 | int32(data[2])<<8 | int32(data[3]), nil
}

func readLong(r *bytes.Reader) (int64, error) {
	data := make([]byte, 8)
	index, err := r.Read(data)
	if err != nil || index != 8 {
		return 0, err
	}
	return int64(data[0])<<56 | int64(data[1])<<48 | int64(data[2])<<40 | int64(data[3])<<32 |
		int64(data[4])<<24 | int64(data[5])<<16 | int64(data[6])<<8 | int64(data[7]), nil
}

func (that *Conn) Receive(packet interface{}) error {
	pak, err := that.ReadPacket()
	if err == nil {
		return ParsePacket(packet, pak)
	}
	return err
}

func ParsePacket(packet interface{}, pak pk.Packet) error {
	type1, value := reflect.TypeOf(packet).Elem(), reflect.ValueOf(packet).Elem()
	field1 := type1.Field(0)
	val, err := strconv.ParseInt(field1.Tag.Get("id"), 10, 8)
	if err != nil {
		return err
	}
	if byte(val) != pak.ID {
		return errors.New("The received packet id (" + string(pak.ID) +
			") is not equal to the needed id (" + string(val) + ")!")
	}
	if field1.Name == "PacketId" {
		return nil
	}
	r := bytes.NewReader(pak.Data)
	for i, num := 0, type1.NumField(); i < num; i++ {
		f, kind := value.Field(i), type1.Field(i).Type
		switch kind.Kind() {
		case reflect.Bool:
			data, err := r.ReadByte()
			if err != nil {
				return err
			}
			f.SetBool(data == 1)
		case reflect.String:
			length, err := readVarInt(r)
			if err != nil {
				return err
			}
			data := make([]byte, length)
			index, err := r.Read(data)
			if err != nil || index != length {
				return err
			}
			f.SetString(string(data))
		case reflect.Uint8:
			data, err := r.ReadByte()
			if err != nil {
				return err
			}
			f.SetUint(uint64(data))
		case reflect.Int8:
			data, err := r.ReadByte()
			if err != nil {
				return err
			}
			f.SetInt(int64(data))
		case reflect.Int16:
			data := make([]byte, 2)
			index, err := r.Read(data)
			if err != nil || index != 2 {
				return err
			}
			f.SetInt(int64(int16(data[0])<<8 | int16(data[1])))
		case reflect.Uint16:
			data := make([]byte, 2)
			index, err := r.Read(data)
			if err != nil || index != 2 {
				return err
			}
			f.SetUint(uint64(int16(data[0])<<8 | int16(data[1])))
		case reflect.Int:
			value, err := readVarInt(r)
			if err != nil {
				return err
			}
			f.SetInt(int64(value))
		case reflect.Int32:
			value, err := readInt(r)
			if err != nil {
				return err
			}
			f.SetInt(int64(value))
		case reflect.Int64:
			if val, ok := type1.Field(i).Tag.Lookup("var"); ok && val == "true" {
				var n uint64
				for i := 0; ; i++ {
					sec, err := r.ReadByte()
					if err != nil {
						return err
					}
					n |= uint64(sec&0x7F) << uint64(7*i)
					if i >= 10 {
						return errors.New("VarLong is too big")
					} else if sec&0x80 == 0 {
						break
					}
				}
				f.SetInt(int64(n))
			} else {
				value, err := readLong(r)
				if err != nil {
					return err
				}
				f.SetInt(value)
			}
		case reflect.Float32:
			value, err := readInt(r)
			if err != nil {
				return err
			}
			f.SetFloat(float64(math.Float32frombits(uint32(value))))
		case reflect.Float64:
			value, err := readLong(r)
			if err != nil {
				return err
			}
			f.SetFloat(math.Float64frombits(uint64(value)))
		case UUID:
			data := uuid.UUID{}
			length, err := io.ReadFull(r, (&data)[:])
			if err != nil || length != 16 {
				return err
			}
			f.Set(reflect.ValueOf(data))
		case POSITION:
			v, err := readLong(r)
			if err != nil {
				return err
			}

			x := int(v >> 38)
			y := int(v & 0xFFF)
			z := int(v << 26 >> 38)

			//处理负数
			if x >= 1<<25 {
				x -= 1 << 26
			}
			if y >= 1<<11 {
				y -= 1 << 12
			}
			if z >= 1<<25 {
				z -= 1 << 26
			}
			f.Set(reflect.ValueOf(pk.Position{X: x, Y: y, Z: z}))
		case reflect.Array:
			if kind.Elem().Kind() == reflect.Uint8 {
				length := r.Len()
				data := make([]byte, length)
				index, err := r.Read(data)
				if err != nil || index != length {
					return err
				}
				f.SetBytes(data)
			}
		}
	}
	return nil
}
