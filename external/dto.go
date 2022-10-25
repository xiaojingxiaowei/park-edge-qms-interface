package external

//========================== req ==============

type DetectReqDTO struct {
	ImageID   string  //图片ID
	ImageLeft float64 //图像物理左边距
	ImageTop  float64 //图像物理顶部边距

	LatitudeIndex  int32 //图像块的纬向标号，从0开始，整数范围的一个数，一般不大于50，表示图片位于第几列
	LongitudeIndex int32 //图像块在经向标号，从0开始，整数范围的一个数，表示图片位于第几行
	LatitudeStatus int32 //0=第一帧图像，1=中间帧图像；2=最后一帧图像

	RotationAngle       int32 //旋转角度，角度值：0-360，以图片的中心为原点旋转
	RotationOrientation int32 //旋转角度方向，0：逆时针，1：顺时针

	Image ImageDataDTO //图像数据

	CreateAt  int64  //创建时间
	Extension string //附加扩展字段（JSON）
}

type ImageDataDTO struct {
	Data      []byte //图像二进制数据
	Width     int32  //width of image
	Height    int32  //height of image
	ImageType int32  //type of image
}

type DetectReqDTO1 struct {
	ImageID   string  //图片ID
	ImageLeft float64 //图像物理左边距
	ImageTop  float64 //图像物理顶部边距

	LatitudeIndex  int32 //图像块的纬向标号，从0开始，整数范围的一个数，一般不大于50，表示图片位于第几列
	LongitudeIndex int32 //图像块在经向标号，从0开始，整数范围的一个数，表示图片位于第几行
	LatitudeStatus int32 //0=第一帧图像，1=中间帧图像；2=最后一帧图像

	RotationAngle       int32 //旋转角度，角度值：0-360，以图片的中心为原点旋转
	RotationOrientation int32 //旋转角度方向，0：逆时针，1：顺时针

	Defects []DefectDTO1 //疵点信息

	CreateAt int64 //创建时间
}

type DefectDTO1 struct {
	DefectID    string //疵点ID
	DefectType  int32  //疵点类型
	DefectCount int32  //疵点数量

	Left   float64 //校正后相对整批位置
	Top    float64 //校正后相对整批位置
	Width  float64 //校正后疵点宽高
	Height float64 //校正后疵点宽高

	Confidence    float64 //置信度(算法提供)，缺陷最下面坐标
	AlarmStatus   int32   //0=不亮灯，1=亮红灯，2=亮绿灯，3=亮黄灯
	SaveStatus    int32
	DefectStatus  int32 //0=正在拼接状态，1=拼接完成状态，其他扩展状态
	DisposeStatus int32 //0=未处理，1=已处理，2=误报

	Image ImageDTO1 //图像数据
}

type ImageDTO1 struct {
	ImageID   string  //图片唯一标识
	ImageLeft float64 //图片在整匹布的坐标，以左上角为坐标原点
	ImageTop  float64 //图片在整匹布的坐标，以左上角为坐标原点

	Left   float64 //瑕疵在图片上的位置
	Top    float64 //瑕疵在图片上的位置
	Width  float64 //图片的宽
	Height float64 //图片的高

	LatitudeIndex  int32 //图像块的纬向标号，从0开始，整数范围的一个数
	LongitudeIndex int32 //图像块在经向标号，从0开始，整数范围的一个数
	LatitudeStatus int32 //图片拼接 - 0=第一帧  1 = 中间帧图像， 2 = 最后一帧图像
}

type InitReqDTO struct {
	OrgCode     string //工厂编码
	VarietyId   int32  //
	EquipmentId string //设备id

	ImageCount int32   //图片数量
	TransRatio float64 //缩放比

	MaxSplit       int32  //最大拼接数
	SmallClothCode string //小卷布票号（小卷质检报告号）
	SmallClothNum  int32  //小卷布匹序列

	RuleConfig  string //规则配置
	AlarmConfig string //告警配置
}

type LoomReqDTO struct {
	Speed      int32  //速度
	Signal     string //信号
	LightColor int32  //亮灯颜色
	Twinkle    int32  //闪灯
}

//========================== res ==============

type DefectInfo struct {
	DefectID   string
	DefectType int32

	Left   float64
	Top    float64
	Width  float64
	Height float64

	Confidence    float64
	AlarmStatus   int32
	SaveStatus    int32
	DefectStatus  int32
	DisposeStatus int32

	SmallClothCode string
	SmallClothNum  int32

	DefectEventInfo DefectEventInfo
	MergeInfos      []MergeInfo
	ImageInfos      []ImageInfo
}

type DefectEventInfo struct {
	Record       int64
	EventType    int32
	Light        int32
	LightTimeout int64
	LightColor   int32
	Position     int32
	TwinkleType  int32
	Weight       int32
}

type MergeInfo struct {
	BeforeMergeDefect []string
	AfterMergeDefect  string
}

type ImageInfo struct {
	ImageID    string
	DefectType int32
	Left       float64
	Top        float64
	Width      float64
	Height     float64

	LatitudeIndex  int32
	LongitudeIndex int32
	LatitudeStatus int32

	ImageLeft  float64
	ImageTop   float64
	Confidence float64

	CurrentImage int32
	Extension    string
}

type VersionInfo struct {
	ModelVersion            string //模型版本
	RegLibVersion           string //算法库版本
	RuleEngineVersion       string //规则引擎包版本
	ApiVersion              string //api版本
	SdkVersion              string //sdk版本
	RuleEngineConfigVersion string //规则引擎配置版本
}
