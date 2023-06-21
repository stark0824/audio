package xerr

// OK 成功返回
const OK uint32 = 200

// UNKNOWN 当从调用其他服务（CP微服务、数据库等）返回错误时，可给客户端返回
const UNKNOWN uint32 = 500

// INVALID_ARGUMENT 无效参数 客户端提交的参数验证未通过
const INVALID_ARGUMENT uint32 = 4220

// NOT_FOUND 未找到对应资源，如查询作品信息但作品已删除时
const NOT_FOUND uint32 = 404

// ALREADY_EXISTS 添加不可重复的数据时，验证到数据已存在，或者用户已经执行过不可重复的操作
const ALREADY_EXISTS uint32 = 1027

// 后台 - 未登录
const NO_LOGIN uint32 = 9996
