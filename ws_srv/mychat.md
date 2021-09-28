```text
# 用户表
ma_im_user {
id,           // ID
re_id,        // 关联ID
re_type,      // 关联类型
avatar,       // 头像
nick_name,    // 昵称
sex,          // 性别
program_id,   // 应用ID
updated_at,   // 更新时间
created_at,   // 创建时间
}

# 用户好友
ma_im_user_friend {
id,           // ID
user_id,      // 用户ID
friend_id,    // 好友ID
status,       // 状态 0=删除 1=正常
program_id,   // 应用ID
updated_at,   // 更新时间
created_at,   // 创建时间
}

# 群组
ma_im_room {
id,           // ID
sn,           // 编号
name,         // 名称
logo,         // 图标
desc,         // 描述
number,       // 成员数（冗余字段）
program_id,   // 应用ID
updated_at,   // 更新时间
created_at,   // 创建时间
}

# 用户群组
ma_im_user_room {
id,           // ID
user_id,      // 用户ID
room_id,      // 群组ID
status,       // 状态 0=删除 1=正常
program_id,   // 应用ID
updated_at,   // 更新时间
created_at,   // 创建时间
}

# 聊天室
ma_im_chat {
id,           // ID
re_id,        // 关联ID
re_type,      // 关联类型 群组（1个）、好友（2个）
new_id,       // 最新消息ID
program_id,   // 应用ID
updated_at,   // 更新时间
created_at,   // 创建时间
}

# 聊天室消息
ma_im_chat_msg {
id,           // ID
uuid,         // 唯一消息标识 (时间戳（毫秒）+ 聊天室ID + 机子编号 + 序号)
chat_id,      // 聊天室ID
type,         // 类型 0=消息 1=图片 2=图文 4=卡片
content,      // 消息内容 JSON
program_id,   // 应用ID
updated_at,   // 更新时间
created_at,   // 创建时间
}

# 用户聊天室
ma_im_user_chat {
id,           // ID
user_id,      // 用户ID
chat_id,      // 聊天室ID
read_id,      // 已阅读消息ID
re_id,        // 关联ID
re_type,      // 关联类型 用户、群组
status,       // 状态 0=删除 1=正常
program_id,   // 应用ID
updated_at,   // 更新时间
created_at,   // 创建时间
}
```