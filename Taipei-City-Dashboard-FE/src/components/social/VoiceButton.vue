<template>
	<div class="voice-control">
		<button @click="startSpeaking">Start Voice Guide</button>
		<button @click="pauseSpeaking">Pause</button>
		<button @click="resumeSpeaking">Resume</button>
		<button @click="cancelSpeaking">Cancel</button>
		<button @click="downloadContent">Download Content</button>
	</div>
</template>

<script setup>
import { ref } from "vue";

let utterance;
const speechSynthesis = window.speechSynthesis;
const isSpeaking = ref(false);
const content = ref("");

const startSpeaking = () => {
	if (!isSpeaking.value) {
		content.value = getSimplifiedContent();
		utterance = new SpeechSynthesisUtterance(content.value);
		utterance.lang = "zh-TW"; // 设置语言，可以根据需要调整
		utterance.rate = 1; // 设置语速
		utterance.pitch = 1; // 设置音调

		utterance.onend = () => {
			isSpeaking.value = false;
		};

		speechSynthesis.speak(utterance);
		isSpeaking.value = true;
	}
};

const pauseSpeaking = () => {
	if (isSpeaking.value && !speechSynthesis.paused) {
		speechSynthesis.pause();
	}
};

const resumeSpeaking = () => {
	if (isSpeaking.value && speechSynthesis.paused) {
		speechSynthesis.resume();
	}
};

const cancelSpeaking = () => {
	if (isSpeaking.value) {
		speechSynthesis.cancel();
		isSpeaking.value = false;
	}
};

const downloadContent = () => {
	const blob = new Blob([content.value], { type: "text/plain" });
	const link = document.createElement("a");
	link.href = URL.createObjectURL(blob);
	link.download = "page-content.txt";
	link.click();
};

const getSimplifiedContent = () => {
	let text = "";
	document.querySelectorAll(".dashboardcomponent").forEach((element) => {
		text += `\n${element.innerText.trim()}\n`;
	});
	return text.replace(/\s+/g, " ").trim();
};
</script>

<style scoped>
.voice-control {
	position: fixed;
	bottom: 100px;
	left: 20px;
	display: flex;
	flex-direction: column;
	gap: 10px;
	background-color: #f8f9fa;
	padding: 10px;
	border-radius: 8px;
	box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
	z-index: 1000;
}

.voice-control button {
	background-color: #007bff;
	color: white;
	padding: 10px;
	border: none;
	border-radius: 5px;
	cursor: pointer;
	transition: background-color 0.3s;
}

.voice-control button:hover {
	background-color: #0056b3;
}
</style>
