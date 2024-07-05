<template>
	<div class="info-button" @click="grabPageInfo">
		<span>Grab Page Info</span>
	</div>
</template>

<script setup>
const grabPageInfo = () => {
	// 选择特定类名的元素
	const elements = document.querySelectorAll(".dashboardcomponent");

	// 遍历 DOM 树并提取信息
	const info = [];
	elements.forEach((element) => {
		const elementInfo = {
			outerHTML: element.outerHTML.trim(), // 包括元素本身及其所有子元素的 HTML 内容
			innerHTML: element.innerHTML.trim(), // 只包括元素的子元素的 HTML 内容
			id: element.id || null,
			class: element.className || null,
		};
		info.push(elementInfo);
	});

	// 将信息转换为 JSON 字符串
	const infoJson = JSON.stringify(info, null, 2);

	// 显示信息或下载

	// 创建隐藏的链接元素以下载 JSON 文件
	const blob = new Blob([infoJson], { type: "application/json" });
	const link = document.createElement("a");
	link.href = URL.createObjectURL(blob);
	link.download = "dashboardcomponent-info.json";
	link.click();
};
</script>

<style scoped>
.info-button {
	position: fixed;
	bottom: 100px;
	left: 20px;
	background-color: #007bff;
	color: white;
	padding: 10px 20px;
	border-radius: 50px;
	cursor: pointer;
	box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
	display: flex;
	align-items: center;
	justify-content: center;
	font-size: 16px;
	transition: background-color 0.3s;
	z-index: 1000;
}

.info-button:hover {
	background-color: #0056b3;
}
</style>
