<template>
	<div class="chat-view">
		<ChatBox />
	</div>
</template>

<script>
import ChatBox from "../components/chat/ChatBox.vue";

export default {
	name: "ChatView",
	components: {
		ChatBox,
	},
};
</script>

<style>
/* 在 ChatBox.vue 的样式中添加以下内容 */
.chat-component {
	width: 100%;
	height: 100%;
	font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
	text-align: center;
	background-color: var(--dashboardcomponent-color-component-background);
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
	border-radius: 5px;
	padding: var(--dashboardcomponent-font-m);
}

.chat {
	width: 100%;
	height: 100%;
	display: flex;
	flex-direction: column;
	justify-content: space-between;
}
</style>
