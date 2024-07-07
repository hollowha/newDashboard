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
					<span class="icon" v-if="message.type === 'announcement'">
						announcement
					</span>
					<span class="icon" v-if="message.type === 'wish'"
						>stars</span
					>
					<span class="icon" v-if="message.type === 'message'"
						>message</span
					>
					<span class="icon" v-if="message.type === 'ai'"
						>smart_toy</span
					>
					<strong>{{ message.username }}:</strong>
					<span v-html="message.message"></span>
				</div>
			</div>
			<div class="input-container">
				<div class="input-row">
					<input v-model="username" placeholder="Username" />
					<div class="select-container">
						<select v-model="messageType" :disabled="isLoading">
							<option value="message">一般訊息</option>
							<option value="announcement">公告訊息</option>
							<option value="wish">許願訊息</option>
							<option value="ai">AI聊天</option>
							<option value="report">異常回報</option>
						</select>
						<div class="icon-overlay">
							<span
								class="icon"
								v-if="messageType === 'announcement'"
							>
								announcement
							</span>
							<span class="icon" v-if="messageType === 'wish'"
								>stars</span
							>
							<span class="icon" v-if="messageType === 'message'"
								>message</span
							>
							<span class="icon" v-if="messageType === 'ai'"
								>smart_toy</span
							>
							<span class="icon" v-if="messageType === 'report'"
								>report</span
							>
						</div>
					</div>
				</div>
				<div v-if="messageType === 'report'">
					<select v-model="resourceType" class="select-container">
						<option value="elec">電力</option>
						<option value="water">水源</option>
						<option value="gas">天然氣</option>
						<option value="road">道路</option>
						<option value="other">其他</option>
					</select>
					<input
						v-model="message"
						class="input-row"
						@keyup.enter="sendMessage"
						placeholder="Message"
						:disabled="isLoading && messageType === 'ai'"
					/>
					<button
						@click="sendMessage"
						class="send-btn"
						:disabled="isLoading && messageType === 'ai'"
					>
						<span class="icon">send</span>
					</button>
				</div>
				<div class="input-row" v-else>
					<input
						v-model="message"
						@keyup.enter="sendMessage"
						placeholder="Message"
						:disabled="isLoading && messageType === 'ai'"
					/>
					<button
						@click="sendMessage"
						class="send-btn"
						:disabled="isLoading && messageType === 'ai'"
					>
						<span class="icon">send</span>
					</button>
				</div>
			</div>
		</div>
		<div v-if="isLoading && messageType === 'ai'" class="loading-overlay">
			<div class="spinner"></div>
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
			conversationHistory: [],
			isLoading: false,
			resourceType: "elec", // 新增的變數
			currentLatLng: "", // 新增的變數
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
	watch: {
		messageType(newValue) {
			if (newValue === "report") {
				this.fetchCurrentLocation();
			}
		},
		resourceType(newValue) {
			this.updateMessage();
		},
	},

	methods: {
		async fetchCurrentLocation() {
			if ("geolocation" in navigator) {
				navigator.geolocation.getCurrentPosition(
					(position) => {
						const { latitude, longitude } = position.coords;
						this.currentLatLng = `${latitude} ${longitude}`;
						this.updateMessage();
						console.log("Current location: ", latitude, longitude);
					},
					(error) => {
						console.error("Geolocation error: ", error);
						this.currentLatLng = "32.2323 4343.4343"; // fallback to fake coordinates
						this.updateMessage();
					}
				);
			} else {
				this.currentLatLng = "32.2323 4343.4343"; // fallback to fake coordinates
				this.updateMessage();
			}
		},
		updateMessage() {
			if (this.messageType === "report") {
				this.message = `detail-here`;
			}
		},
		async sendMessage() {
			if (this.isLoading && this.messageType === "ai") return;
			if (this.username && this.message) {
				let formattedMessage = this.message;
				if (this.messageType === "announcement") {
					formattedMessage = "!a " + this.message;
				} else if (this.messageType === "wish") {
					formattedMessage = "!w " + this.message;
				} else if (this.messageType === "report") {
					formattedMessage = `!no ${this.resourceType} ${this.currentLatLng} ${this.message}`;
				}

				const msg = {
					username: this.username,
					message: formattedMessage,
					type: this.messageType,
					dashboardDisplay: this.dashboardIndex,
				};

				if (this.messageType !== "ai") {
					this.ws.send(JSON.stringify(msg));
				} else {
					this.isLoading = true;
					this.addToConversationHistory("user", formattedMessage);
					this.displayMessage(formattedMessage, this.username);
					const aiResponse = await this.sendAIMessage(
						formattedMessage
					);
					this.addToConversationHistory("ai", aiResponse);
					this.displayMessage(aiResponse, "AI");
					this.isLoading = false;
				}
				this.message = "";
			}
		},
		// 其他方法保持不變

		// 其他方法保持不變

		addToConversationHistory(role, text) {
			this.conversationHistory.push({
				role: role,
				parts: [
					{
						text: text,
					},
				],
			});
		},
		async sendAIMessage(message) {
			try {
				const response = await fetch(
					"https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=AIzaSyBfyF26G-RvZPlfMSSuDbUtjHNOFqYnk8Y",
					{
						method: "POST",
						headers: {
							"Content-Type": "application/json",
						},
						body: JSON.stringify({
							contents: [
								{
									parts: [
										{
											text: `不要用markdown格式，請用純文字。用簡短精簡文字，
											臺北城市儀表板是由臺北大數據中心（TUIC）開發的資料視覺化平台。
我們的主要目標是建立一個全方位的資料視覺化工具，協助台北市政府的治理決策，相較其他類似的資料視覺化平台，臺北城市儀表板以其全面而簡約直觀的設計脫穎而出。它的建立旨在鼓勵跨時間和空間的交叉參照，將資料集組織成視覺化的模組組件，並透過視覺化帶來更多不同以往的發現與洞見。
臺北城市儀表板顯示的大部分數據來自台北的開放資料平台data.taipei。我們更為儀表板上的每個資料集提供數據解釋指南和建議的使用案例，讓任何用戶能充分探索台北開放資料的種種可能性。
臺北城市儀表板是第一個開源的城市儀表板。我們不僅公開儀表板的程式碼，也邀請使用者一起加入開發的行列。提供的資訊與功能有: 臺北市陳情系統, 臺北市陳情系統, 刑事統計, 犯罪點位, 警察相關設施, 緊急避難設施, 火災風險地點, 火災應變設施, 山坡地風險地點, 道路施工, 最新交通事故月統計, 公車行駛狀態, 都市更新案件, 社會住宅建設進度, 閒置市有財產, 居職人口推估, 公有土地, 都市計畫用地類型, 全市屋齡分布, 抽水站狀態, 交通路況, 水位監測, YouBike使用情況, 松山新店線擁擠程度, 派工案件, 捷運人流趨勢, 板南線擁擠程度, 社福機構, 緊急救護服務, 社福人口, 公車站, 公園綠地, 都市綠地, 人行道分布, 醫療院所分布, AED分布, 路燈分佈, 自行車道, 行道樹, 田園城市, 公共停車場, 土壤液化潛勢, 淡水信義線擁擠程度, 中和新蘆線擁擠程度, 建照使照發照量, 垃圾車收運點位, YouBike週末群像, YouBike週間群像, YouBike見車率, YouBike互補站點, YouBike設站狀態, YouBike未設站區潛在需求, 育兒支持資源, 育嬰假性別比, 居住安全性, 婦產科診所與醫院分佈, 幼兒人數統計, 哺集乳室分佈, 生育補助合約醫院, 防洪與雨水設備, 透水鋪面, 降雨淹水模擬圖, 危樓面積, 防空避難設施, 狹小巷弄程度分佈, 溫室氣體排放統計, 空氣品質, 用電量統計, 監視器分佈, 道路管制, 腸病毒流行情形警示
在這次黑客松新增的功能有: 留言區, 許願池, 公告, 截圖, 分享, 追蹤儀錶板, 熱門排行, GPTs, 移動到當前位置, 地點問題回報, 組件地圖預覽, 反戰資料展示
											請你簡短回答這次訊息: ${message} 之前對話紀錄如下：${JSON.stringify(
												this.conversationHistory
											)}




											`,
										},
									],
								},
							],
						}),
					}
				);
				const data = await response.json();
				return data.candidates[0].content.parts[0].text;
			} catch (error) {
				console.error("Error:", error);
				return "AI 服務無法使用。";
			}
		},
		displayMessage(message, sender) {
			const messageElement = {
				username: sender,
				message: message,
				type: sender === "AI" ? "ai" : this.messageType,
			};
			this.messages.push(messageElement);
		},
	},
};
// !no elec 32.2323 4343.4343 auto generate
</script>

<style scoped>
.icon {
	font-family: var(--dashboardcomponent-font-icon);
	padding: var(--font-s);
	font-weight: normal !important;
}

.chat-component {
	position: relative;
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
	background-color: #8c763f; /* 顏色可根據需要調整 */
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

.loading-overlay {
	position: absolute;
	top: 0;
	left: 0;
	width: 100%;
	height: 100%;
	background-color: rgba(0, 0, 0, 0.5);
	display: flex;
	justify-content: center;
	align-items: center;
	z-index: 10;
}

.spinner {
	border: 4px solid rgba(255, 255, 255, 0.3);
	border-top: 4px solid #fff;
	border-radius: 50%;
	width: 40px;
	height: 40px;
	animation: spin 1s linear infinite;
}

@keyframes spin {
	to {
		transform: rotate(360deg);
	}
}
</style>
