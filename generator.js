const data = require('./minecraft-data/data/pc/1.16.1/protocol.json')

const mapping = {
    varint: 'int',
    varlong: 'int64',
    buffer: '[]byte',
    restBuffer: '[]byte',
    entityMetadata: '[]byte',
    optionalNbt: '[]byte',
    slot: '[]byte',
    nbt: 'packet.NBT',
    u8: 'byte',
    u16: 'uint16',
    i8: 'int8',
    i16: 'int16',
    i32: 'int32',
    i64: 'int64',
    f32: 'float32',
    f64: 'float64',
    UUID: 'uuid.UUID',
    position: 'packet.Position'
}

const firstUpperCase = str => str[0].toUpperCase() + str.slice(1)
const toUpperCase = str => str.replace(/_(\w)/g, (_, it) => it.toUpperCase())
let output = `package go_mc_wrapped

import (
    "github.com/Tnze/go-mc/net/packet"
    "github.com/google/uuid"
)\n\n`
delete data.types
const append = (type, kind, json) => {
    let i = 0
    for (const k in json) {
        if (!k.startsWith('packet_')) return
        output += `type ${type}${toUpperCase(k.replace('packet', ''))}To${kind} struct {\n`
        if (json[k][1].length === 0) output += '    PacketId []byte `id:"' + i +'"`\n'
        else json[k][1].forEach((it, index) => {
            if (!it.name) {
                output += '    UnknownField []byte\n'
                return
            }
            output += `    ${toUpperCase(firstUpperCase(it.name))} `
            switch (it.type[0]) {
                case 'option':
                    it.type.shift()
                    if (it.type.length === 1) it.type = it.type[0]
                    break;
                case 'switch':
                case 'array':
                case 'particleData':
                case 'topBitSetTerminatedArray':
                    it.type = 'buffer'
            }
            let skip = false
            if (it.type in mapping) output += mapping[it.type]
            else if (it.type[0] === 'buffer' && it.type[1] && it.type[1].countType === 'varint') {
                output += '[]byte `count:"true"`'
                skip = true
            } else output += it.type
            if (!index && !skip) output += ' `id:"' + i + '"`'
            output += '\n'
        })
        output += "}\n\n"
        i++
    }
}
for (const k in data) {
    const v = data[k]
    const name = firstUpperCase(k)
    append(name, 'Server', v.toServer.types)
    append(name, 'Client', v.toClient.types)
}
require('fs').writeFileSync('types.go', output)
