<!-- Developed by Taipei Urban Intelligence Center 2023-2024-->

<script setup>
import { onMounted, ref } from "vue";
import { useContentStore } from "../../../store/contentStore";
import { useDialogStore } from "../../../store/dialogStore";
import { useMapStore } from "../../../store/mapStore";
import { useAuthStore } from "../../../store/authStore";
// import { useRouter } from "vue-router"; // 引入 useRouter
// import ChatBox from "../../../views/ChatView.vue"; // 引入 ChatBox 组件
// import ChatBox from "../../chat/ChatBox.vue"; // 引入 ChatBox 组件

import SideBarTab from "../miscellaneous/SideBarTab.vue";

const contentStore = useContentStore();
const dialogStore = useDialogStore();
const mapStore = useMapStore();
const authStore = useAuthStore();
// const router = useRouter(); // 初始化 useRouter
// The expanded state is also stored in localstorage to retain the setting after refresh
const isExpanded = ref(true);
// const showChatBox = ref(false); // 控制 ChatBox 显示的状态

function handleOpenAddDashboard() {
	dialogStore.addEdit = "add";
	dialogStore.showDialog("addEditDashboards");
}

function toggleExpand() {
	isExpanded.value = isExpanded.value ? false : true;
	localStorage.setItem("isExpanded", isExpanded.value);
	if (!isExpanded.value) {
		mapStore.resizeMap();
	}
}

// function toggleChatBox() {
// 	showChatBox.value = !showChatBox.value; // 切换 ChatBox 的显示状态
// }

onMounted(() => {
	const storedExpandedState = localStorage.getItem("isExpanded");
	if (storedExpandedState === "false") {
		isExpanded.value = false;
	} else {
		isExpanded.value = true;
	}
});
</script>

<template>
	<div
		:class="{
			sidebar: true,
			'sidebar-collapse': !isExpanded,
			'hide-if-mobile': true,
		}"
	>
		<div v-if="authStore.token">
			<h2>{{ isExpanded ? `我的最愛` : `最愛` }}</h2>
			<SideBarTab
				icon="favorite"
				title="收藏組件"
				:expanded="isExpanded"
				:index="contentStore.favorites?.index"
			/>
			<div class="sidebar-sub-add">
				<h2>{{ isExpanded ? `個人儀表板 ` : `個人` }}</h2>
				<button v-if="isExpanded" @click="handleOpenAddDashboard">
					<span>add_circle_outline</span>新增
				</button>
			</div>
			<div
				v-if="
					contentStore.personalDashboards.filter(
						(item) => item.icon !== 'favorite'
					).length === 0
				"
				class="sidebar-sub-no"
			>
				<p>{{ isExpanded ? `尚無個人儀表板 ` : `尚無` }}</p>
			</div>
			<SideBarTab
				v-for="item in contentStore.personalDashboards.filter(
					(item) => item.icon !== 'favorite'
				)"
				:key="item.index"
				:icon="item.icon"
				:title="item.name"
				:index="item.index"
				:expanded="isExpanded"
			/>
		</div>
		<h2>{{ isExpanded ? `公共儀表板 ` : `公共` }}</h2>
		<SideBarTab
			v-for="item in contentStore.publicDashboards.filter(
				(item) => item.index !== 'map-layers'
			)"
			:key="item.index"
			:icon="item.icon"
			:title="item.name"
			:index="item.index"
			:expanded="isExpanded"
		/>

		<h2>{{ isExpanded ? `基本地圖圖層` : `圖層` }}</h2>
		<SideBarTab
			icon="public"
			title="圖資資訊"
			:expanded="isExpanded"
			index="map-layers"
		/>
		<!-- <button @click="toggleChatBox" class="chat">
			<h2>匿名聊天</h2>
			<span class="chat-span">chat</span>
		</button> -->

		<!-- 在這裡添加點擊事件來導航到 /chat -->
		<!-- <button @click="navigateToChat" class="chat">
			<h2>匿名聊天</h2>
			<span class="chat-span">chat</span>
		</button> -->

		<button class="sidebar-collapse-button" @click="toggleExpand">
			<span>{{
				isExpanded
					? "keyboard_double_arrow_left"
					: "keyboard_double_arrow_right"
			}}</span>
		</button>
	</div>

	<!-- 浮动 ChatBox -->
	<!-- <div v-if="showChatBox" class="floating-chatbox">
		<ChatBox />
	</div> -->
</template>

<style scoped lang="scss">
.chat {
	max-height: var(--font-xl);
	display: flex;
	align-items: center;
	margin: var(--font-s) 0;
	border-left: solid 4px transparent;
	border-radius: 0 5px 5px 0;
	transition: background-color 0.2s;
	white-space: nowrap;
	text-wrap: nowrap;
}
.chat-span {
	font-family: var(--dashboardcomponent-font-icon);
	padding: var(--font-s);
}
.sidebar {
	width: 170px;
	min-width: 170px;
	height: calc(100vh - 80px);
	height: calc(var(--vh) * 100 - 80px);
	max-height: calc(100vh - 80px);
	max-height: calc(var(--vh) * 100 - 80px);
	position: relative;
	padding: 0 10px 0 var(--font-m);
	margin-top: 20px;
	border-right: 1px solid var(--color-border);
	transition: min-width 0.2s ease-out;
	overflow-x: hidden;
	overflow-y: scroll;
	user-select: none;

	h2 {
		color: var(--color-complement-text);
		font-weight: 400;
		text-wrap: nowrap;
	}

	&-sub {
		margin-bottom: var(--font-s);

		&-add {
			width: 100%;
			display: flex;
			text-wrap: nowrap;

			button {
				display: flex;
				align-items: center;
				flex-wrap: nowrap;
				margin-left: 0.5rem;
				padding: 2px 6px;
				border-radius: 5px;
				background-color: var(--color-highlight);
				color: var(--color-normal-text);
				text-wrap: nowrap;

				span {
					margin-right: 4px;
					font-family: var(--font-icon);
				}
			}
		}

		&-no p {
			margin: 0.5rem 0 0.5rem 10px;
			font-size: var(--font-s);
			font-style: italic;
		}
	}

	&-collapse {
		width: 45px;
		min-width: 45px;

		h2 {
			margin-left: 5px;
		}

		&-button {
			height: fit-content;
			position: absolute;
			bottom: 10px;
			right: 10px;
			padding: 5px;
			border-radius: 5px;
			transition: background-color 0.2s;

			&:hover {
				background-color: var(--color-component-background);
			}

			span {
				font-family: var(--font-icon);
				font-size: var(--font-l);
			}
		}
	}
}

.floating-chatbox {
	position: fixed;
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
	width: 100%; /* 设置浮动窗体的宽度 */
	height: 400px; /* 设置浮动窗体的高度 */
	box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1); /* 阴影效果 */
	border-radius: 10px; /* 圆角效果 */
	z-index: 1000; /* 确保浮动窗体显示在最上层 */
	border-radius: 5px;
}
</style>
