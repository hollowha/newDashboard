<!-- Developed By Taipei Urban Intelligence Center 2023-2024 -->
<!-- 
Lead Developer:  Igor Ho (Full Stack Engineer)
Data Pipelines:  Iima Yu (Data Scientist)
Design and UX: Roy Lin (Fmr. Consultant), Chu Chen (Researcher)
Systems: Ann Shih (Systems Engineer)
Testing: Jack Huang (Data Scientist), Ian Huang (Data Analysis Intern) 
-->
<!-- Department of Information Technology, Taipei City Government -->

<script setup>
import { onBeforeMount, onMounted, onBeforeUnmount, ref, computed } from "vue";
import { useAuthStore } from "./store/authStore";
import { useDialogStore } from "./store/dialogStore";
import { useContentStore } from "./store/contentStore";

import NavBar from "./components/utilities/bars/NavBar.vue";
import SideBar from "./components/utilities/bars/SideBar.vue";
import AdminSideBar from "./components/utilities/bars/AdminSideBar.vue";
import SettingsBar from "./components/utilities/bars/SettingsBar.vue";
import NotificationBar from "./components/dialogs/NotificationBar.vue";
import InitialWarning from "./components/dialogs/InitialWarning.vue";
import ComponentSideBar from "./components/utilities/bars/ComponentSideBar.vue";
import LogIn from "./components/dialogs/LogIn.vue";
import ShareButton from "./components/social/ShareButton.vue";
import ScreenshotButton from "./components/social/ScreenshotButton.vue";
import ChatBox from "./components/chat/ChatBox.vue";
// import VoiceButton from "./components/social/VoiceButton.vue";

const authStore = useAuthStore();
const dialogStore = useDialogStore();
const contentStore = useContentStore();

const timeToUpdate = ref(600);
const showChatBox = ref(false); // 控制 ChatBox 显示的状态
const showButtons = ref(false); // 控制分享和截图按钮的显示状态

const formattedTimeToUpdate = computed(() => {
	const minutes = Math.floor(timeToUpdate.value / 60);
	const seconds = timeToUpdate.value % 60;
	return `${minutes}:${seconds < 10 ? "0" : ""}${seconds}`;
});

function reloadChartData() {
	if (!["dashboard", "mapview"].includes(authStore.currentPath)) return;
	contentStore.setCurrentDashboardChartData();
	timeToUpdate.value = 600;
}
function updateTimeToUpdate() {
	if (!["dashboard", "mapview"].includes(authStore.currentPath)) return;
	if (timeToUpdate.value <= 0) {
		timeToUpdate.value = 0;
		return;
	}
	timeToUpdate.value -= 5;
}

function toggleChatBox() {
	showChatBox.value = !showChatBox.value; // 切换 ChatBox 的显示状态
}

function toggleButtons() {
	showButtons.value = !showButtons.value; // 切换分享和截图按钮的显示状态
}
onBeforeMount(() => {
	authStore.initialChecks();

	let vh = window.innerHeight * 0.01;
	document.documentElement.style.setProperty("--vh", `${vh}px`);

	window.addEventListener("resize", () => {
		let vh = window.innerHeight * 0.01;
		document.documentElement.style.setProperty("--vh", `${vh}px`);
	});
});
onMounted(() => {
	const showInitialWarning = localStorage.getItem("initialWarning");
	if (!showInitialWarning && authStore.currentPath !== "embed") {
		dialogStore.showDialog("initialWarning");
	}

	setInterval(reloadChartData, 1000 * 600);
	setInterval(updateTimeToUpdate, 1000 * 5);
});
onBeforeUnmount(() => {
	clearInterval(reloadChartData);
	clearInterval(updateTimeToUpdate);
});
</script>

<template>
	<div class="app-container">
		<NotificationBar />
		<NavBar v-if="authStore.currentPath !== 'embed'" />
		<!-- /mapview, /dashboard layouts -->
		<div
			v-if="
				authStore.currentPath === 'mapview' ||
				authStore.currentPath === 'dashboard'
			"
			class="app-content"
		>
			<SideBar />
			<div class="app-content-main">
				<SettingsBar />
				<RouterView />
			</div>
		</div>
		<!-- /admin layouts -->
		<div v-else-if="authStore.currentPath === 'admin'" class="app-content">
			<AdminSideBar />
			<div class="app-content-main">
				<RouterView />
			</div>
		</div>
		<!-- /component, /component/:index layouts -->
		<div
			v-else-if="authStore.currentPath.includes('component')"
			class="app-content"
		>
			<ComponentSideBar />
			<div class="app-content-main">
				<RouterView />
			</div>
		</div>
		<div v-else>
			<router-view />
		</div>
		<InitialWarning />
		<LogIn />
		<div
			v-if="
				['dashboard', 'mapview'].includes(authStore.currentPath) &&
				!authStore.isMobile &&
				!authStore.isNarrowDevice
			"
			class="app-update"
		>
			<p>下次更新：{{ formattedTimeToUpdate }}</p>
		</div>
		<!-- <ShareButton />
		<ScreenshotButton />

		<button @click="toggleChatBox" class="chat">
			<h2>匿名聊天</h2>
			<span class="chat-span">chat</span>
		</button> -->
		<!-- 浮動 ChatBox -->
		<div class="button-container">
			<button @click="toggleButtons" class="main-button">
				<h2>聊天&分享</h2>
				<span class="icon">menu</span>
			</button>
			<div v-if="showButtons" class="secondary-buttons">
				<button @click="toggleChatBox" class="secondary-button">
					<span class="icon">chat</span>
				</button>
				<ScreenshotButton class="secondary-button" />
				<ShareButton class="secondary-button" />
			</div>
		</div>
		<div v-if="showChatBox" class="floating-chatbox">
			<ChatBox />
		</div>
	</div>
</template>

<style scoped lang="scss">
.app {
	&-container {
		max-width: 100vw;
		max-height: 100vh;
		max-height: calc(var(--vh) * 100);
	}

	&-content {
		width: 100vw;
		max-width: 100vw;
		height: calc(100vh - 60px);
		height: calc(var(--vh) * 100 - 60px);
		display: flex;

		&-main {
			width: 100%;
			display: flex;
			flex-direction: column;
		}
	}

	&-update {
		position: fixed;
		bottom: 0;
		right: 20px;
		color: white;
		opacity: 0.3;
		transition: opacity 0.3s;
		user-select: none;

		p {
			color: var(--color-complement-text);
		}

		&:hover {
			opacity: 1;
		}
	}
}

.button-container {
	position: fixed;
	bottom: 20px;
	right: 20px;
	display: flex;
	flex-direction: column;
	align-items: center;
	z-index: 1000;

	.main-button {
		background-color: var(--color-highlight);
		color: white;
		padding: 10px;
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
		box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
		transition: background-color 0.3s;
		cursor: pointer;
		width: 60px;
		height: 60px;
		border: 2px solid white; /* 白色边框 */

		h2 {
			display: none;
		}

		&:hover {
			background-color: lightblue; /* 浅蓝色 */
		}

		.icon {
			font-family: var(--font-icon);
			font-size: var(--font-xl);
		}
	}

	.secondary-buttons {
		display: flex;
		flex-direction: column;
		margin-top: 10px;

		.secondary-button {
			background-color: var(--color-component-background);
			color: white;
			padding: 10px;
			border-radius: 50%;
			display: flex;
			align-items: center;
			justify-content: center;
			box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
			transition: background-color 0.3s;
			cursor: pointer;
			width: 50px;
			height: 50px;
			margin-bottom: 10px;
			border: 2px solid white; /* 白色边框 */

			&:hover {
				background-color: black; /* 深黑色 */
			}

			.icon {
				font-family: var(--font-icon);
				font-size: var(--font-l);
			}
		}
	}
}

.floating-chatbox {
	position: fixed;
	bottom: 20px;
	right: 90px;
	width: 400px; /* 設定浮動視窗的寬度 */
	height: auto; /* 設定浮動視窗的高度 */
	box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1); /* 陰影效果 */
	border-radius: 10px; /* 圓角效果 */
	z-index: 1000; /* 確保浮動視窗顯示在最上層 */
	overflow-y: scroll;
	border: 2px solid #ffffff;
	background-color: var(--color-component-background);
}
</style>
