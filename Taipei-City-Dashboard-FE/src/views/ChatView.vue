<template>
	<div>
		<!-- <h1>Chat View</h1> -->
		<ChatBox />
	</div>
	<!-- <DashboardComponent
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
		/> -->
	<!-- <DashboardComponent
		v-for="item in contentStore.currentDashboard.components"
		v-if="item.index === 'pump_status'"
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
			authStore.token && contentStore.currentDashboard.icon !== 'favorite'
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
	/> -->
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
