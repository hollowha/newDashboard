<template>
	<div class="chat-component">
		<div class="chat">
			<h1>WebSocket Chat</h1>
			<div class="messages">
				<div
					v-for="(message, index) in messages"
					:key="index"
					:class="['message', message.type]"
				>
					<span class="icon" v-if="message.type !== 'message'">{{
						message.type === "announcement"
							? "announcement"
							: "stars"
					}}</span>
					<strong>{{ message.username }}:</strong>
					{{ message.message }}
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
import { useAuthStore } from "../../store/authStore.js";

export default {
	props: {
		dashboardIndex: {
			type: String,
			required: false,
			default: "0",
		},
	},
	data() {
		return {
			ws: null,
			username: "",
			message: "",
			messages: [],
		};
	},
	created() {
		const authStore = useAuthStore();
		if (authStore.user && authStore.user.name) {
			this.username = authStore.user.name;
		}

		this.ws = new WebSocket("ws://localhost:8088/ws");
		this.ws.onmessage = (event) => {
			const msg = JSON.parse(event.data);
			if (msg.dashboardDisplay === this.dashboardIndex) {
				this.messages.push(msg);
			}
		};
	},
	methods: {
		sendMessage() {
			if (this.username && this.message) {
				const msg = {
					username: this.username,
					message: this.message,
					dashboardDisplay: this.dashboardIndex,
				};
				this.ws.send(JSON.stringify(msg));
				this.message = "";
			}
		},
		// get the geojson data from the public/mapData folder
		// getGeoJson(){


		// }
	},
};
</script>

<style scoped>
.icon {
	font-family: var(--dashboardcomponent-font-icon);
	padding: var(--font-s);
	font-weight: normal !important;
}

.chat-component {
	font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
	color: #fff;
	text-align: center;
	background-color: var(--dashboardcomponent-color-component-background);
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
	border-radius: 5px;
	padding: var(--dashboardcomponent-font-m);
}

h1 {
	font-size: 1.5em; /* 調整標題大小，使其在較小的高度下更適合 */
	margin-bottom: 10px;
	color: #fff;
}

.chat {
	height: 370px;
	width: 100%;
	/*max-width: 600px;*/
	display: flex;
	flex-direction: column;
	height: 100%;
	justify-content: space-between;
}

.messages {
	border: 1px solid #555;
	flex-grow: 1;
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
	border-radius: 4px;
	color: #fff;
}

.message.announcement {
	background-color: #228176; /* 顏色可根據需要調整 */
	font-size: 1.2em;
	font-weight: bold;
}

.message.wish {
	background-color: #8c763f; /* 顏色可根據需要調整 */
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
