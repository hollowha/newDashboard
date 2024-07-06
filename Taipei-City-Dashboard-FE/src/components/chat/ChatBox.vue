<template>
	<div class="chat-component">
		<div class="chat">
			<h1>聊天分享區</h1>
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
				<div class="input-row">
					<input v-model="username" placeholder="Username" />
					<div class="select-container">
						<select v-model="messageType">
							<option value="message">一般訊息</option>
							<option value="announcement">公告訊息</option>
							<option value="wish">許願訊息</option>
						</select>
						<div class="icon-overlay">
							<span
								class="icon"
								v-if="messageType === 'announcement'"
								>announcement</span
							>
							<span class="icon" v-if="messageType === 'wish'"
								>stars</span
							>
							<span class="icon" v-if="messageType === 'message'"
								>message</span
							>
						</div>
					</div>
				</div>
				<div class="input-row">
					<input
						v-model="message"
						@keyup.enter="sendMessage"
						placeholder="Message"
					/>
					<button @click="sendMessage" class="send-btn">
						<span class="icon">send</span>
					</button>
				</div>
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
			messageType: "message",
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
				let formattedMessage = this.message;
				if (this.messageType === "announcement") {
					formattedMessage = "!a " + this.message;
				} else if (this.messageType === "wish") {
					formattedMessage = "!w " + this.message;
				}
				const msg = {
					username: this.username,
					message: formattedMessage,
					type: this.messageType,
					dashboardDisplay: this.dashboardIndex,
				};
				this.ws.send(JSON.stringify(msg));
				this.message = "";
			}
		},
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
	overflow-y: scroll;
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
	overflow-y: scroll;
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
	background-color: #8c763f; /* 顏色可根據需要調整 */ /* 顏色可根據需要調整 */
	font-size: 1.2em;
	font-weight: bold;
}

.message.wish {
	background-color: #228176;
}

.input-container {
	display: flex;
	flex-direction: column;
	gap: 10px;
	width: 100%;
}

.input-row {
	display: flex;
	flex-direction: row;
	gap: 10px;
	width: 100%;
}

input,
select {
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
	color: #fff;
	cursor: pointer;
	transition: background 0.3s ease;
}

button:hover {
	color: #85f9e7;
}

.send-btn:hover {
	color: #85f9e7;
}

.select-container {
	position: relative;
	flex-grow: 1;
	display: flex;
	align-items: center;
}

.icon-overlay {
	position: absolute;
	top: 50%;
	left: 10px;
	transform: translateY(-50%);
	pointer-events: none;
	display: flex;
	align-items: center;
	z-index: 2;
}

select {
	-webkit-appearance: none;
	-moz-appearance: none;
	appearance: none;
	width: 100%;
	padding-left: 40px; /* Adjust padding to leave space for the icon */
	position: relative;
	z-index: 1;
	background-color: #222;
	color: #fff;
	border: 1px solid #555;
	border-radius: 4px;
}
</style>
