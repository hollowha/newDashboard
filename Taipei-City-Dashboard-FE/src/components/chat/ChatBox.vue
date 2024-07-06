<template>
	<div id="chat-component">
		<h1>WebSocket Chat</h1>
		<div class="chat">
			<div class="messages">
				<div
					v-for="(message, index) in messages"
					:key="index"
					class="message"
				>
					<strong>{{ message.username }}:</strong>
					{{ message.message }}
					{{ message }}
				</div>
			</div>
			<div class="input-container">
				<input v-model="username" placeholder="Username" />
				<input
					v-model="message"
					@keyup.enter="sendMessage"
					placeholder="Message"
				/>
				<button @click="sendMessage">Send</button>
			</div>
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

<style scoped>
#chat-component {
	font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
	background: #000;
	color: #fff;
	text-align: center;
	padding: 20px;
	height: 100vh;
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
}

h1 {
	font-size: 2em;
	margin-bottom: 20px;
	color: #fff;
}

.chat {
	background: #333;
	border-radius: 8px;
	box-shadow: 0 2px 10px rgba(0, 0, 0, 0.5);
	width: 100%;
	max-width: 600px;
	padding: 20px;
}

.messages {
	border: 1px solid #555;
	height: 300px;
	overflow-y: auto;
	margin-bottom: 10px;
	padding: 10px;
	border-radius: 4px;
	background: #444;
	color: #fff;
}

.message {
	margin-bottom: 10px;
	padding: 5px;
	background: #555;
	border-radius: 4px;
	color: #fff;
}

.input-container {
	display: flex;
	flex-direction: row;
	gap: 10px;
}

input {
	padding: 10px;
	border-radius: 4px;
	border: 1px solid #555;
	outline: none;
	flex-grow: 1;
	font-size: 1em;
	background: #222;
	color: #fff;
}

input::placeholder {
	color: #bbb;
}

button {
	padding: 10px 20px;
	border: none;
	border-radius: 4px;
	background: #555;
	color: #fff;
	font-size: 1em;
	cursor: pointer;
	transition: background 0.3s ease;
}

button:hover {
	background: #777;
}
</style>
