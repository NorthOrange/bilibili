// 资源访问路径
package tools

func AvatarPath(avatarName string) string { // 将头像名转换为可访问的头像路径
	return Config.Socket + "/avatar/" + avatarName
}

func VideoPath(videoPath string) string { // 将视频存储路径转换为可访问路径
	return Config.Socket + "/video/" + videoPath
}
