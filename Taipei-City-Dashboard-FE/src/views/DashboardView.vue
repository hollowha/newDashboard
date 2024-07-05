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
import { DashboardComponent } from "city-dashboard-component";
import { useContentStore } from "../store/contentStore";
import { useDialogStore } from "../store/dialogStore";
import { useAuthStore } from "../store/authStore";

import MoreInfo from "../components/dialogs/MoreInfo.vue";
import ReportIssue from "../components/dialogs/ReportIssue.vue";

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

// function toggleFavorite(id) {
//   if (contentStore.favorites.components.includes(id)) {
//     // 呼叫 API 來取消收藏
//     axios.post('/api/unfavorite', { componentId: id })
//       .then(response => {
//         console.log('Unfavorited successfully:', response);
//         contentStore.unfavoriteComponent(id);
//       })
//       .catch(error => {
//         console.error('Error unfavoriting component:', error);
//       });
//   } else {
//     // 呼叫 API 來添加收藏
//     axios.post('/api/favorite', { componentId: id })
//       .then(response => {
//         console.log('Favorited successfully:', response);
//         contentStore.favoriteComponent(id);
//       })
//       .catch(error => {
//         console.error('Error favoriting component:', error);
//       });
//   }
// }

function toggleFavorite(id) {
	if (contentStore.favorites.components.includes(id)) {
		// 模擬取消收藏的 API 調用
		console.log(`Unfavoriting component with id: ${id}`);
		// 模擬成功回應
		setTimeout(() => {
			console.log("Unfavorited successfully");
			contentStore.unfavoriteComponent(id);
		}, 250); // 模擬網絡延遲
	} else {
		// 模擬添加收藏的 API 調用
		console.log(`Favoriting component with id: ${id}`);
		// 模擬成功回應
		setTimeout(() => {
			console.log("Favorited successfully");
			contentStore.favoriteComponent(id);
		}, 250); // 模擬網絡延遲
	}
}
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
		<div>test_like</div>

		<MoreInfo />
		<ReportIssue />
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
