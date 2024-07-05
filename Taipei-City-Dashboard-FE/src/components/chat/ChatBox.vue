<template>
	<div id="app">
		<h1>WebSocket Chat</h1>
		<div class="chat">
			<div class="messages">
				<div v-for="(message, index) in messages" :key="index">
					<strong>{{ message.username }}:</strong>
					{{ message.message }}
				</div>
			</div>
			<input v-model="username" placeholder="Username" />
			<input
				v-model="message"
				@keyup.enter="sendMessage"
				placeholder="Message"
			/>
			<button @click="sendMessage">Send</button>
		</div>
	</div>
</template>

<script>
export default {
	data() {
		return {
			ws: null,
			username: "",
			message: "",
			messages: [],
		};
	},
	created() {
		this.ws = new WebSocket("ws://localhost:8088/ws");
		this.ws.onmessage = (event) => {
			const msg = JSON.parse(event.data);
			this.messages.push(msg);
		};
	},
	methods: {
		sendMessage() {
			if (this.username && this.message) {
				const msg = {
					username: this.username,
					message: this.message,
				};
				this.ws.send(JSON.stringify(msg));
				this.message = "";
			}
		},
	},
};
</script>

<style>
.chat {
	max-width: 600px;
	margin: 0 auto;
}
.messages {
	border: 1px solid #ccc;
	height: 300px;
	overflow-y: scroll;
	margin-bottom: 10px;
	padding: 10px;
}
input {
	margin-right: 5px;
}
</style>
