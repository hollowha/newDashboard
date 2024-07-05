<template>
	<div class="screenshot-button" @click="downloadScreenshot">
		<span>Download Screenshot</span>
	</div>
</template>

<script setup>
import html2canvas from "html2canvas";

const downloadScreenshot = async () => {
	try {
		const canvas = await html2canvas(document.body);
		const dataUrl = canvas.toDataURL("image/png");
		const link = document.createElement("a");
		link.href = dataUrl;
		link.download = "screenshot.png";
		link.click();
	} catch (error) {
		console.error("Error taking screenshot:", error);
	}
};
</script>

<style scoped>
.screenshot-button {
	position: fixed;
	bottom: 60px;
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

.screenshot-button:hover {
	background-color: #0056b3;
}
</style>
