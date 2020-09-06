package go_mc_wrapped

import (
	"github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
)

type HandshakingSetProtocolToServer struct {
	ProtocolVersion int `id:"0"`
	ServerHost      string
	ServerPort      uint16
	NextState       int
}

type HandshakingLegacyServerListPingToServer struct {
	Payload byte `id:"1"`
}

type StatusPingStartToServer struct {
	PacketId []byte `id:"0"`
}

type StatusPingToServer struct {
	Time int64 `id:"1"`
}

type StatusServerInfoToClient struct {
	Response string `id:"0"`
}

type StatusPingToClient struct {
	Time int64 `id:"1"`
}

type LoginLoginStartToServer struct {
	Username string `id:"0"`
}

type LoginEncryptionBeginToServer struct {
	SharedSecret []byte `count:"true"`
	VerifyToken  []byte `count:"true"`
}

type LoginLoginPluginResponseToServer struct {
	MessageId int `id:"2"`
	Data      []byte
}

type LoginDisconnectToClient struct {
	Reason string `id:"0"`
}

type LoginEncryptionBeginToClient struct {
	ServerId    string `id:"1"`
	PublicKey   []byte `count:"true"`
	VerifyToken []byte `count:"true"`
}

type LoginSuccessToClient struct {
	Uuid     uuid.UUID `id:"2"`
	Username string
}

type LoginCompressToClient struct {
	Threshold int `id:"3"`
}

type LoginLoginPluginRequestToClient struct {
	MessageId int `id:"4"`
	Channel   string
	Data      []byte
}

type PlayTeleportConfirmToServer struct {
	TeleportId int `id:"0"`
}

type PlayQueryBlockNbtToServer struct {
	TransactionId int `id:"1"`
	Location      packet.Position
}

type PlaySetDifficultyToServer struct {
	NewDifficulty byte `id:"2"`
}

type PlayEditBookToServer struct {
	NewBook []byte `id:"3"`
	Signing bool
	Hand    int
}

type PlayQueryEntityNbtToServer struct {
	TransactionId int `id:"4"`
	EntityId      int
}

type PlayPickItemToServer struct {
	Slot int `id:"5"`
}

type PlayNameItemToServer struct {
	Name string `id:"6"`
}

type PlaySelectTradeToServer struct {
	Slot int `id:"7"`
}

type PlaySetBeaconEffectToServer struct {
	PrimaryEffect   int `id:"8"`
	SecondaryEffect int
}

type PlayUpdateCommandBlockToServer struct {
	Location packet.Position `id:"9"`
	Command  string
	Mode     int
	Flags    byte
}

type PlayUpdateCommandBlockMinecartToServer struct {
	EntityId    int `id:"10"`
	Command     string
	TrackOutput bool
}

type PlayUpdateStructureBlockToServer struct {
	Location  packet.Position `id:"11"`
	Action    int
	Mode      int
	Name      string
	OffsetX   byte
	OffsetY   byte
	OffsetZ   byte
	SizeX     byte
	SizeY     byte
	SizeZ     byte
	Mirror    int
	Rotation  int
	Metadata  string
	Integrity float32
	Seed      int
	Flags     byte
}

type PlayTabCompleteToServer struct {
	TransactionId int `id:"12"`
	Text          string
}

type PlayChatToServer struct {
	Message string `id:"13"`
}

type PlayClientCommandToServer struct {
	ActionId int `id:"14"`
}

type PlaySettingsToServer struct {
	Locale       string `id:"15"`
	ViewDistance int8
	ChatFlags    int
	ChatColors   bool
	SkinParts    byte
	MainHand     int
}

type PlayTransactionToServer struct {
	WindowId int8 `id:"16"`
	Action   int16
	Accepted bool
}

type PlayEnchantItemToServer struct {
	WindowId    int8 `id:"17"`
	Enchantment int8
}

type PlayWindowClickToServer struct {
	WindowId    byte `id:"18"`
	Slot        int16
	MouseButton int8
	Action      int16
	Mode        int8
	Item        []byte
}

type PlayCloseWindowToServer struct {
	WindowId byte `id:"19"`
}

type PlayCustomPayloadToServer struct {
	Channel string `id:"20"`
	Data    []byte
}

type PlayUseEntityToServer struct {
	Target   int `id:"21"`
	Mouse    int
	X        []byte
	Y        []byte
	Z        []byte
	Hand     []byte
	Sneaking bool
}

type PlayGenerateStructureToServer struct {
	Location    packet.Position `id:"22"`
	Levels      int
	KeepJigsaws bool
}

type PlayKeepAliveToServer struct {
	KeepAliveId int64 `id:"23"`
}

type PlayLockDifficultyToServer struct {
	Locked bool `id:"24"`
}

type PlayPositionToServer struct {
	X        float64 `id:"25"`
	Y        float64
	Z        float64
	OnGround bool
}

type PlayPositionLookToServer struct {
	X        float64 `id:"26"`
	Y        float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
}

type PlayLookToServer struct {
	Yaw      float32 `id:"27"`
	Pitch    float32
	OnGround bool
}

type PlayFlyingToServer struct {
	OnGround bool `id:"28"`
}

type PlayVehicleMoveToServer struct {
	X     float64 `id:"29"`
	Y     float64
	Z     float64
	Yaw   float32
	Pitch float32
}

type PlaySteerBoatToServer struct {
	LeftPaddle  bool `id:"30"`
	RightPaddle bool
}

type PlayCraftRecipeRequestToServer struct {
	WindowId int8 `id:"31"`
	Recipe   string
	MakeAll  bool
}

type PlayAbilitiesToServer struct {
	Flags int8 `id:"32"`
}

type PlayBlockDigToServer struct {
	Status   int8 `id:"33"`
	Location packet.Position
	Face     int8
}

type PlayEntityActionToServer struct {
	EntityId  int `id:"34"`
	ActionId  int
	JumpBoost int
}

type PlaySteerVehicleToServer struct {
	Sideways float32 `id:"35"`
	Forward  float32
	Jump     byte
}

type PlayCraftingBookDataToServer struct {
	Type         int `id:"36"`
	UnknownField []byte
}

type PlayResourcePackReceiveToServer struct {
	Result int `id:"37"`
}

type PlayHeldItemSlotToServer struct {
	SlotId int16 `id:"38"`
}

type PlaySetCreativeSlotToServer struct {
	Slot int16 `id:"39"`
	Item []byte
}

type PlayUpdateJigsawBlockToServer struct {
	Location   packet.Position `id:"40"`
	Name       string
	Target     string
	Pool       string
	FinalState string
	JointType  string
}

type PlayUpdateSignToServer struct {
	Location packet.Position `id:"41"`
	Text1    string
	Text2    string
	Text3    string
	Text4    string
}

type PlayArmAnimationToServer struct {
	Hand int `id:"42"`
}

type PlaySpectateToServer struct {
	Target uuid.UUID `id:"43"`
}

type PlayBlockPlaceToServer struct {
	Hand        int `id:"44"`
	Location    packet.Position
	Direction   int
	CursorX     float32
	CursorY     float32
	CursorZ     float32
	InsideBlock bool
}

type PlayUseItemToServer struct {
	Hand int `id:"45"`
}

type PlayAdvancementTabToServer struct {
	Action int `id:"46"`
	TabId  []byte
}

type PlaySpawnEntityToClient struct {
	EntityId   int `id:"0"`
	ObjectUUID uuid.UUID
	Type       int
	X          float64
	Y          float64
	Z          float64
	Pitch      int8
	Yaw        int8
	ObjectData int32
	VelocityX  int16
	VelocityY  int16
	VelocityZ  int16
}

type PlaySpawnEntityExperienceOrbToClient struct {
	EntityId int `id:"1"`
	X        float64
	Y        float64
	Z        float64
	Count    int16
}

type PlaySpawnEntityLivingToClient struct {
	EntityId   int `id:"2"`
	EntityUUID uuid.UUID
	Type       int
	X          float64
	Y          float64
	Z          float64
	Yaw        int8
	Pitch      int8
	HeadPitch  int8
	VelocityX  int16
	VelocityY  int16
	VelocityZ  int16
}

type PlaySpawnEntityPaintingToClient struct {
	EntityId   int `id:"3"`
	EntityUUID uuid.UUID
	Title      int
	Location   packet.Position
	Direction  byte
}

type PlayNamedEntitySpawnToClient struct {
	EntityId   int `id:"4"`
	PlayerUUID uuid.UUID
	X          float64
	Y          float64
	Z          float64
	Yaw        int8
	Pitch      int8
}

type PlayAnimationToClient struct {
	EntityId  int `id:"5"`
	Animation byte
}

type PlayStatisticsToClient struct {
	Entries []byte `id:"6"`
}

type PlayAdvancementsToClient struct {
	Reset              bool `id:"7"`
	AdvancementMapping []byte
	Identifiers        []byte
	ProgressMapping    []byte
}

type PlayBlockBreakAnimationToClient struct {
	EntityId     int `id:"8"`
	Location     packet.Position
	DestroyStage int8
}

type PlayTileEntityDataToClient struct {
	Location packet.Position `id:"9"`
	Action   byte
	NbtData  []byte
}

type PlayBlockActionToClient struct {
	Location packet.Position `id:"10"`
	Byte1    byte
	Byte2    byte
	BlockId  int
}

type PlayBlockChangeToClient struct {
	Location packet.Position `id:"11"`
	Type     int
}

type PlayBossBarToClient struct {
	EntityUUID uuid.UUID `id:"12"`
	Action     int
	Title      []byte
	Health     []byte
	Color      []byte
	Dividers   []byte
	Flags      []byte
}

type PlayDifficultyToClient struct {
	Difficulty       byte `id:"13"`
	DifficultyLocked bool
}

type PlayTabCompleteToClient struct {
	TransactionId int `id:"14"`
	Start         int
	Length        int
	Matches       []byte
}

type PlayDeclareCommandsToClient struct {
	Nodes     []byte `id:"15"`
	RootIndex int
}

type PlayFacePlayerToClient struct {
	FeetEyes       int `id:"16"`
	X              float64
	Y              float64
	Z              float64
	IsEntity       bool
	EntityId       []byte
	EntityFeetEyes []byte
}

type PlayNbtQueryResponseToClient struct {
	TransactionId int `id:"17"`
	Nbt           packet.NBT
}

type PlayChatToClient struct {
	Message  string `id:"18"`
	Position int8
	Sender   uuid.UUID
}

type PlayMultiBlockChangeToClient struct {
	ChunkX  int32 `id:"19"`
	ChunkZ  int32
	Records []byte
}

type PlayTransactionToClient struct {
	WindowId int8 `id:"20"`
	Action   int16
	Accepted bool
}

type PlayCloseWindowToClient struct {
	WindowId byte `id:"21"`
}

type PlayOpenWindowToClient struct {
	WindowId      int `id:"22"`
	InventoryType int
	WindowTitle   string
}

type PlayWindowItemsToClient struct {
	WindowId byte `id:"23"`
	Items    []byte
}

type PlayCraftProgressBarToClient struct {
	WindowId byte `id:"24"`
	Property int16
	Value    int16
}

type PlaySetSlotToClient struct {
	WindowId int8 `id:"25"`
	Slot     int16
	Item     []byte
}

type PlaySetCooldownToClient struct {
	ItemID        int `id:"26"`
	CooldownTicks int
}

type PlayCustomPayloadToClient struct {
	Channel string `id:"27"`
	Data    []byte
}

type PlayNamedSoundEffectToClient struct {
	SoundName     string `id:"28"`
	SoundCategory int
	X             int32
	Y             int32
	Z             int32
	Volume        float32
	Pitch         float32
}

type PlayKickDisconnectToClient struct {
	Reason string `id:"29"`
}

type PlayEntityStatusToClient struct {
	EntityId     int32 `id:"30"`
	EntityStatus int8
}

type PlayExplosionToClient struct {
	X                    float32 `id:"31"`
	Y                    float32
	Z                    float32
	Radius               float32
	AffectedBlockOffsets []byte
	PlayerMotionX        float32
	PlayerMotionY        float32
	PlayerMotionZ        float32
}

type PlayUnloadChunkToClient struct {
	ChunkX int32 `id:"32"`
	ChunkZ int32
}

type PlayGameStateChangeToClient struct {
	Reason   byte `id:"33"`
	GameMode float32
}

type PlayOpenHorseWindowToClient struct {
	WindowId byte `id:"34"`
	NbSlots  int
	EntityId int32
}

type PlayKeepAliveToClient struct {
	KeepAliveId int64 `id:"35"`
}

type PlayMapChunkToClient struct {
	X             int32 `id:"36"`
	Z             int32
	GroundUp      bool
	IgnoreOldData bool
	BitMap        int
	Heightmaps    packet.NBT
	Biomes        []byte
	ChunkData     []byte `count:"true"`
	BlockEntities []byte
}

type PlayWorldEventToClient struct {
	EffectId int32 `id:"37"`
	Location packet.Position
	Data     int32
	Global   bool
}

type PlayWorldParticlesToClient struct {
	ParticleId   int32 `id:"38"`
	LongDistance bool
	X            float64
	Y            float64
	Z            float64
	OffsetX      float32
	OffsetY      float32
	OffsetZ      float32
	ParticleData float32
	Particles    int32
	Data         []byte
}

type PlayUpdateLightToClient struct {
	ChunkX              int `id:"39"`
	ChunkZ              int
	TrustEdges          bool
	SkyLightMask        int
	BlockLightMask      int
	EmptySkyLightMask   int
	EmptyBlockLightMask int
	Data                []byte
}

type PlayLoginToClient struct {
	EntityId            int32 `id:"40"`
	GameMode            byte
	PreviousGameMode    byte
	WorldNames          []byte
	DimensionCodec      packet.NBT
	Dimension           string
	WorldName           string
	HashedSeed          int64
	MaxPlayers          byte
	ViewDistance        int
	ReducedDebugInfo    bool
	EnableRespawnScreen bool
	IsDebug             bool
	IsFlat              bool
}

type PlayMapToClient struct {
	ItemDamage       int `id:"41"`
	Scale            int8
	TrackingPosition bool
	Locked           bool
	Icons            []byte
	Columns          int8
	Rows             []byte
	X                []byte
	Y                []byte
	Data             []byte
}

type PlayTradeListToClient struct {
	WindowId          int `id:"42"`
	Trades            []byte
	VillagerLevel     int
	Experience        int
	IsRegularVillager bool
	CanRestock        bool
}

type PlayRelEntityMoveToClient struct {
	EntityId int `id:"43"`
	DX       int16
	DY       int16
	DZ       int16
	OnGround bool
}

type PlayEntityMoveLookToClient struct {
	EntityId int `id:"44"`
	DX       int16
	DY       int16
	DZ       int16
	Yaw      int8
	Pitch    int8
	OnGround bool
}

type PlayEntityLookToClient struct {
	EntityId int `id:"45"`
	Yaw      int8
	Pitch    int8
	OnGround bool
}

type PlayEntityToClient struct {
	EntityId int `id:"46"`
}

type PlayVehicleMoveToClient struct {
	X     float64 `id:"47"`
	Y     float64
	Z     float64
	Yaw   float32
	Pitch float32
}

type PlayOpenBookToClient struct {
	Hand int `id:"48"`
}

type PlayOpenSignEntityToClient struct {
	Location packet.Position `id:"49"`
}

type PlayCraftRecipeResponseToClient struct {
	WindowId int8 `id:"50"`
	Recipe   string
}

type PlayAbilitiesToClient struct {
	Flags        int8 `id:"51"`
	FlyingSpeed  float32
	WalkingSpeed float32
}

type PlayCombatEventToClient struct {
	Event    int `id:"52"`
	Duration []byte
	PlayerId []byte
	EntityId []byte
	Message  []byte
}

type PlayPlayerInfoToClient struct {
	Action int `id:"53"`
	Data   []byte
}

type PlayPositionToClient struct {
	X          float64 `id:"54"`
	Y          float64
	Z          float64
	Yaw        float32
	Pitch      float32
	Flags      int8
	TeleportId int
}

type PlayUnlockRecipesToClient struct {
	Action             int `id:"55"`
	CraftingBookOpen   bool
	FilteringCraftable bool
	SmeltingBookOpen   bool
	FilteringSmeltable bool
	Recipes1           []byte
	Recipes2           []byte
}

type PlayEntityDestroyToClient struct {
	EntityIds []byte `id:"56"`
}

type PlayRemoveEntityEffectToClient struct {
	EntityId int `id:"57"`
	EffectId int8
}

type PlayResourcePackSendToClient struct {
	Url  string `id:"58"`
	Hash string
}

type PlayRespawnToClient struct {
	Dimension        string `id:"59"`
	WorldName        string
	HashedSeed       int64
	Gamemode         byte
	PreviousGamemode byte
	IsDebug          bool
	IsFlat           bool
	CopyMetadata     bool
}

type PlayEntityHeadRotationToClient struct {
	EntityId int `id:"60"`
	HeadYaw  int8
}

type PlayWorldBorderToClient struct {
	Action         int `id:"61"`
	Radius         []byte
	X              []byte
	Z              []byte
	OldRadius      []byte
	NewRadius      []byte
	Speed          []byte
	PortalBoundary []byte
	WarningTime    []byte
	WarningBlocks  []byte
}

type PlayCameraToClient struct {
	CameraId int `id:"62"`
}

type PlayHeldItemSlotToClient struct {
	Slot int8 `id:"63"`
}

type PlayUpdateViewPositionToClient struct {
	ChunkX int `id:"64"`
	ChunkZ int
}

type PlayUpdateViewDistanceToClient struct {
	ViewDistance int `id:"65"`
}

type PlayScoreboardDisplayObjectiveToClient struct {
	Position int8 `id:"66"`
	Name     string
}

type PlayEntityMetadataToClient struct {
	EntityId int `id:"67"`
	Metadata []byte
}

type PlayAttachEntityToClient struct {
	EntityId  int32 `id:"68"`
	VehicleId int32
}

type PlayEntityVelocityToClient struct {
	EntityId  int `id:"69"`
	VelocityX int16
	VelocityY int16
	VelocityZ int16
}

type PlayEntityEquipmentToClient struct {
	EntityId   int `id:"70"`
	Equipments []byte
}

type PlayExperienceToClient struct {
	ExperienceBar   float32 `id:"71"`
	Level           int
	TotalExperience int
}

type PlayUpdateHealthToClient struct {
	Health         float32 `id:"72"`
	Food           int
	FoodSaturation float32
}

type PlayScoreboardObjectiveToClient struct {
	Name        string `id:"73"`
	Action      int8
	DisplayText []byte
	Type        []byte
}

type PlaySetPassengersToClient struct {
	EntityId   int `id:"74"`
	Passengers []byte
}

type PlayTeamsToClient struct {
	Team              string `id:"75"`
	Mode              int8
	Name              []byte
	FriendlyFire      []byte
	NameTagVisibility []byte
	CollisionRule     []byte
	Formatting        []byte
	Prefix            []byte
	Suffix            []byte
	Players           []byte
}

type PlayScoreboardScoreToClient struct {
	ItemName  string `id:"76"`
	Action    int8
	ScoreName string
	Value     []byte
}

type PlaySpawnPositionToClient struct {
	Location packet.Position `id:"77"`
}

type PlayUpdateTimeToClient struct {
	Age  int64 `id:"78"`
	Time int64
}

type PlayTitleToClient struct {
	Action  int `id:"79"`
	Text    []byte
	FadeIn  []byte
	Stay    []byte
	FadeOut []byte
}

type PlayEntitySoundEffectToClient struct {
	SoundId       int `id:"80"`
	SoundCategory int
	EntityId      int
	Volume        float32
	Pitch         float32
}

type PlayStopSoundToClient struct {
	Flags  int8 `id:"81"`
	Source []byte
	Sound  []byte
}

type PlaySoundEffectToClient struct {
	SoundId       int `id:"82"`
	SoundCategory int
	X             int32
	Y             int32
	Z             int32
	Volume        float32
	Pitch         float32
}

type PlayPlayerlistHeaderToClient struct {
	Header string `id:"83"`
	Footer string
}

type PlayCollectToClient struct {
	CollectedEntityId int `id:"84"`
	CollectorEntityId int
	PickupItemCount   int
}

type PlayEntityTeleportToClient struct {
	EntityId int `id:"85"`
	X        float64
	Y        float64
	Z        float64
	Yaw      int8
	Pitch    int8
	OnGround bool
}

type PlayEntityUpdateAttributesToClient struct {
	EntityId   int `id:"86"`
	Properties []byte
}

type PlayEntityEffectToClient struct {
	EntityId      int `id:"87"`
	EffectId      int8
	Amplifier     int8
	Duration      int
	HideParticles int8
}

type PlaySelectAdvancementTabToClient struct {
	Id string `id:"88"`
}

type PlayDeclareRecipesToClient struct {
	Recipes []byte `id:"89"`
}
