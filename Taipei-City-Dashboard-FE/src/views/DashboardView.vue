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
import axios from "axios";
import { DashboardComponent } from "city-dashboard-component";
import { useContentStore } from "../store/contentStore";
import { useDialogStore } from "../store/dialogStore";
import { useAuthStore } from "../store/authStore";
import ChatBox from "../components/chat/ChatBox.vue";

import MoreInfo from "../components/dialogs/MoreInfo.vue";
import ReportIssue from "../components/dialogs/ReportIssue.vue";
import { ref } from "vue";
import { onMounted } from "vue";

const contentStore = useContentStore();
const dialogStore = useDialogStore();
const authStore = useAuthStore();

function handleOpenSettings() {
	contentStore.editDashboard = JSON.parse(
		JSON.stringify(contentStore.currentDashboard)
	);
	dialogStore.addEdit = "edit";
	dialogStore.showDialog("addEditDashboards");
}

// function toggleFavorite(id) {
// 	if (contentStore.favorites.components.includes(id)) {
// 		contentStore.unfavoriteComponent(id);
// 	} else {
// 		contentStore.favoriteComponent(id);
// 	}
// }

// async function toggleFavorite(id) {
// 	try {
// 		// const jwtToken = authStore.token; // 假设 JWT 令牌存储在 authStore 中
// 		const formData = new FormData();
// 		formData.append("componentid", id);

// 		// Authorization: `Bearer ${jwtToken}`,

// 		const response = await axios.get(
// 			`http://localhost:8088/api/v1/like/${id}`,
// 			{
// 				headers: {
// 					"Content-Type": "multipart/form-data",
// 				},
// 				data: formData,
// 			}
// 		);

// 		if (response.status === 200) {
// 			if (contentStore.favorites.components.includes(id)) {
// 				contentStore.unfavoriteComponent(id);
// 			} else {
// 				contentStore.favoriteComponent(id);
// 			}
// 		} else {
// 			console.error(
// 				"Failed to toggle favorite:",
// 				response.status,
// 				response.statusText
// 			);
// 		}
// 	} catch (error) {
// 		console.error("Error toggling favorite:", error);
// 	}
// }

// 定義按鈕狀態
const isFavorited = ref(false);

// 切換按鈕狀態的函數
// 切換按鈕狀態的函數
// const toggleFavorite = async (id) => {
// 	try {
// 		const jwtToken = authStore.token; // 假设 JWT 令牌存储在 authStore 中
// 		console.log(`Toggling like for component id: ${id}`); // 调试信息
// 		const response = await axios.post(
// 			`http://localhost:8088/api/v1/like/${id}`,
// 			{ componentid: id },
// 			{
// 				headers: {
// 					Authorization: `Bearer ${jwtToken}`,
// 					"Content-Type": "application/json",
// 				},
// 			}
// 		);
// 		console.log("Response:", response.data); // 输出返回值
// 		if (response.status === 200) {
// 			isFavorited.value = !isFavorited.value;
// 			console.log(`Successfully toggled like for component id: ${id}`); // 调试信息
// 		} else {
// 			console.error(
// 				"Failed to toggle like:",
// 				response.status,
// 				response.statusText
// 			);
// 		}
// 	} catch (error) {
// 		console.error("Error toggling like:", error);
// 	}
// };

// 定義按鈕狀態

// 检查当前是否已点赞的函数
const checkIfLiked = async (id) => {
	try {
		const jwtToken = authStore.token; // 假設 JWT 令牌存儲在 authStore 中
		const response = await axios.get(
			`http://localhost:8088/api/v1/like/is-like/${id}`,
			{
				headers: {
					Authorization: `Bearer ${jwtToken}`,
					"Content-Type": "application/json",
				},
			}
		);
		if (response.status === 200) {
			isFavorited.value = response.data.is_liked;
			// 更新本地收藏狀態
			if (response.data.is_liked) {
				contentStore.favoriteComponent(id);
			} else {
				contentStore.unfavoriteComponent(id);
			}
		} else {
			console.error(
				"Failed to check if liked:",
				response.status,
				response.statusText
			);
		}
	} catch (error) {
		console.error("Error checking if liked:", error);
	}
};

const toggleFavorite = async (id) => {
	try {
		const jwtToken = authStore.token; // 假設 JWT 令牌存儲在 authStore 中
		console.log(`Toggling like for component id: ${id}`); // 調試信息
		const response = await axios.post(
			`http://localhost:8088/api/v1/like/${id}`,
			{ componentid: id },
			{
				headers: {
					Authorization: `Bearer ${jwtToken}`,
					"Content-Type": "application/json",
				},
			}
		);
		console.log("Response:", response.data); // 輸出返回值
		if (response.status === 200) {
			isFavorited.value = response.data.is_liked;
			console.log(`Successfully toggled like for component id: ${id}`); // 調試信息
			// 更新本地收藏狀態
			if (response.data.is_liked) {
				contentStore.favoriteComponent(id);
			} else {
				contentStore.unfavoriteComponent(id);
			}
		} else {
			console.error(
				"Failed to toggle like:",
				response.status,
				response.statusText
			);
		}
	} catch (error) {
		console.error("Error toggling like:", error);
	}
};

onMounted(() => {
	if (contentStore.currentDashboard.components) {
		contentStore.currentDashboard.components.forEach((component) => {
			checkIfLiked(component.id);
		});
	}
});

// GET http://localhost:8088/api/v1/like/:componentid
// GET
</script>

<template>
	<!-- 1. If the dashboard is map-layers -->
	<div
		v-if="contentStore.currentDashboard.index === 'map-layers'"
		class="dashboard"
	>
		<DashboardComponent
			v-for="item in contentStore.currentDashboard.components"
			:key="item.index"
			v-like-button="item.id"
			:config="item"
			mode="half"
			:info-btn="true"
			:favorite-btn="authStore.token ? true : false"
			:is-favorite="contentStore.favorites?.components.includes(item.id)"
			@favorite="
				(id) => {
					toggleFavorite(id);
				}
			"
			@info="
				(item) => {
					dialogStore.showMoreInfo(item);
				}
			"
		/>
		<MoreInfo />
		<ReportIssue />
	</div>
	<!-- 2. Dashboards that have components -->
	<div
		v-else-if="contentStore.currentDashboard.components?.length !== 0"
		class="dashboard"
	>
		<DashboardComponent
			v-for="item in contentStore.currentDashboard.components"
			:key="item.index"
			v-like-button="item.id"
			:config="item"
			:info-btn="true"
			:delete-btn="
				contentStore.personalDashboards
					.map((item) => item.index)
					.includes(contentStore.currentDashboard.index)
			"
			:favorite-btn="
				authStore.token &&
				contentStore.currentDashboard.icon !== 'favorite'
			"
			:is-favorite="contentStore.favorites?.components.includes(item.id)"
			@favorite="
				(id) => {
					toggleFavorite(id);
				}
			"
			@info="
				(item) => {
					dialogStore.showMoreInfo(item);
				}
			"
			@delete="
				(id) => {
					contentStore.deleteComponent(id);
				}
			"
		/>

		<MoreInfo />
		<ReportIssue />
		<ChatBox />
	</div>

	<!-- <div
		v-else-if="contentStore.currentDashboard.components?.length !== 0"
		class="dashboard"
	>
		<div
			v-for="item in contentStore.currentDashboard.components"
			:key="item.index"
			v-like-button="item.id"
			class="wrapper"
		>
			<DashboardComponent
				:config="item"
				:info-btn="true"
				:delete-btn="
					contentStore.personalDashboards
						.map((item) => item.index)
						.includes(contentStore.currentDashboard.index)
				"
				:favorite-btn="
					authStore.token &&
					contentStore.currentDashboard.icon !== 'favorite'
				"
				:is-favorite="
					contentStore.favorites?.components.includes(item.id)
				"
				@favorite="(id) => toggleFavorite(id)"
				@info="(item) => dialogStore.showMoreInfo(item)"
				@delete="(id) => contentStore.deleteComponent(id)"
			/>
		</div>
		<MoreInfo />
		<ReportIssue />
	</div> -->

	<!-- 3. If dashboard is still loading -->
	<div
		v-else-if="contentStore.loading"
		class="dashboard dashboard-nodashboard"
	>
		<div class="dashboard-nodashboard-content">
			<div />
		</div>
	</div>
	<!-- 4. If dashboard failed to load -->
	<div v-else-if="contentStore.error" class="dashboard dashboard-nodashboard">
		<div class="dashboard-nodashboard-content">
			<span>sentiment_very_dissatisfied</span>
			<h2>發生錯誤，無法載入儀表板</h2>
		</div>
	</div>
	<!-- 5. Dashboards that don't have components -->
	<div v-else class="dashboard dashboard-nodashboard">
		<div class="dashboard-nodashboard-content">
			<span>addchart</span>
			<h2>尚未加入組件</h2>
			<button
				v-if="contentStore.currentDashboard.icon !== 'favorite'"
				class="hide-if-mobile"
				@click="handleOpenSettings"
			>
				加入您的第一個組件
			</button>
			<p v-else>點擊其他儀表板組件之愛心以新增至收藏組件</p>
		</div>
	</div>
</template>

<style scoped lang="scss">
.dashboard {
	max-height: calc(100vh - 127px);
	max-height: calc(var(--vh) * 100 - 127px);
	display: grid;
	row-gap: var(--font-s);
	column-gap: var(--font-s);
	margin: var(--font-m) var(--font-m);
	overflow-y: scroll;

	@media (min-width: 720px) {
		grid-template-columns: 1fr 1fr;
	}

	@media (min-width: 1200px) {
		grid-template-columns: 1fr 1fr 1fr;
	}

	@media (min-width: 1800px) {
		grid-template-columns: 1fr 1fr 1fr 1fr;
	}

	@media (min-width: 2200px) {
		grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
	}

	&-nodashboard {
		grid-template-columns: 1fr;

		&-content {
			width: 100%;
			height: calc(100vh - 127px);
			height: calc(var(--vh) * 100 - 127px);
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;

			span {
				margin-bottom: var(--font-ms);
				font-family: var(--font-icon);
				font-size: 2rem;
			}

			button {
				color: var(--color-highlight);
			}

			div {
				width: 2rem;
				height: 2rem;
				border-radius: 50%;
				border: solid 4px var(--color-border);
				border-top: solid 4px var(--color-highlight);
				animation: spin 0.7s ease-in-out infinite;
			}
		}
	}
}

@keyframes spin {
	to {
		transform: rotate(360deg);
	}
}

.wrapper {
	position: relative;
	display: inline-block; /* 確保不改變內部佈局 */
}

.like-button {
	position: absolute;
	right: 10px;
	bottom: 10px;
	z-index: 10;
	background-color: #ff4081;
	color: white;
	border: none;
	border-radius: 5px;
	padding: 5px 10px;
	cursor: pointer;
}

.dashboardcomponent {
	position: relative;
	width: 500px;
	height: 500px;
}
</style>
