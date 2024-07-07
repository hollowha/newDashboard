<template>
	<div class="report-form">
		<h1>報告資源缺乏位置</h1>
		<form @submit.prevent="submitForm">
			<div class="form-group">
				<label for="theType">資源類型</label>
				<select v-model="form.theType" id="theType" required>
					<option value="elec">電力</option>
					<option value="water">水源</option>
					<option value="gas">天然氣</option>
					<option value="road">道路</option>
					<option value="other">其他</option>
				</select>
			</div>
			<div class="form-group">
				<label for="lat">緯度</label>
				<input
					type="number"
					v-model="form.lat"
					id="lat"
					required
					step="any"
				/>
			</div>
			<div class="form-group">
				<label for="lng">經度</label>
				<input
					type="number"
					v-model="form.lng"
					id="lng"
					required
					step="any"
				/>
			</div>
			<div class="form-group">
				<label for="message">消息</label>
				<textarea
					v-model="form.message"
					id="message"
					required
				></textarea>
			</div>
			<button type="submit">提交</button>
		</form>
		<div v-if="responseMessage" class="response-message">
			{{ responseMessage }}
		</div>
		<!-- <button @click="props">關閉</button> -->
	</div>
</template>

<script>
import axios from "axios";
import { ref, onMounted } from "vue";

export default {
	props: {
		onClose: {
			type: Function,
			required: true,
		},
	},
	setup() {
		const form = ref({
			theType: "elec",
			lat: "",
			lng: "",
			message: "",
		});
		const responseMessage = ref("");

		onMounted(() => {
			if ("geolocation" in navigator) {
				navigator.geolocation.getCurrentPosition(
					(position) => {
						const { latitude, longitude } = position.coords;
						form.value.lat = latitude;
						form.value.lng = longitude;
					},
					(error) => {
						console.error("Geolocation error: ", error);
					}
				);
			}
		});

		async function submitForm() {
			try {
				const response = await axios.post(
					"http://localhost:8088/api/dev/noresource",
					{
						theType: form.value.theType,
						lat: form.value.lat,
						lng: form.value.lng,
						message: form.value.message,
					}
				);
				responseMessage.value = response.data.message || "報告成功！";
			} catch (error) {
				responseMessage.value = "報告失敗，請重試。";
				console.error("Error:", error);
			}
		}

		return {
			form,
			responseMessage,
			submitForm,
		};
	},
};
</script>

<style scoped>
.report-form {
	font-family: Arial, sans-serif;
	max-width: 500px;
	margin: 0 auto;
	padding: 20px;
	background-color: #3e3e3e; /* 灰底 */
	border: 2px solid #ccc; /* 白框 */
	border-radius: 8px;
	box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

h1 {
	text-align: center;
	color: #ffffff;
}

.form-group {
	margin-bottom: 15px;
}

label {
	display: block;
	margin-bottom: 5px;
	color: #ffffff;
}

input,
select,
textarea {
	width: 95%;
	border: 1px solid #ccc;
	border-radius: 4px;
	font-size: 16px;
	background-color: #000000; /* 黑底 */
	color: white; /* 白字 */
}

button {
	display: block;
	width: 100%;
	padding: 10px;
	background-color: var(--color-highlight);
	color: white;
	border: none;
	border-radius: 4px;
	font-size: 16px;
	cursor: pointer;
	transition: background-color 0.3s ease;
}

button:hover {
	background-color: #0056b3;
}

.response-message {
	margin-top: 20px;
	text-align: center;
	color: #333;
}
</style>
