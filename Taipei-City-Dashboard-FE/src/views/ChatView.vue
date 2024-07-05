<template>
	<div>
		<h1>Chat View</h1>
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
/* Add your custom styles here */
</style>
