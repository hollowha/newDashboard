<!-- Developed by Taipei Urban Intelligence Center 2023-2024-->

<!-- This component has two modes "expanded" and "collapsed" which is controlled by the prop "expanded" -->

<script setup>
import { ref, computed } from "vue";
import { useRoute } from "vue-router";
import { useAuthStore } from "../../../store/authStore";
import axios from "axios";

const route = useRoute();

const props = defineProps({
	icon: { type: String },
	title: { type: String },
	index: { type: String },
	expanded: { type: Boolean },
});

const authStore = useAuthStore();

const tabLink = computed(() => {
	if (authStore.currentPath === "admin") {
		return `/admin/${props.index}`;
	}
	return `${route.path}?index=${props.index}`;
});
const linkActiveOrNot = computed(() => {
	if (authStore.currentPath === "admin") {
		return route.path === `/admin/${props.index}` ? true : false;
	}
	return route.query.index === props.index ? true : false;
});

// 定義按鈕狀態
const isFavorited = ref(false);

// 切換按鈕狀態的函數
const toggleFavorite = async () => {
	try {
		const jwtToken = authStore.token; // 假设 JWT 令牌存储在 authStore 中
		const response = await axios.post(
			`http://localhost:8088/api/v1/follow/${props.index}`,
			null,
			{
				headers: {
					Authorization: `Bearer ${jwtToken}`,
					"Content-Type": "application/json",
				},
			}
		);
		if (response.status === 200) {
			isFavorited.value = !isFavorited.value;
		} else {
			console.error(
				"Failed to toggle favorite:",
				response.status,
				response.statusText
			);
		}
	} catch (error) {
		console.error("Error toggling favorite:", error);
	}
};
// 跳轉到指定URL的函數
const goToURL = () => {
	window.location.href =
		"https://chatgpt.com/g/g-IvGrrKzg0-tai-bei-zhan-zheng-fang-wei-yi-biao-ban-zhu-li";
};
</script>

<template>
	<div class="sidebar-item">
		<router-link
			:to="tabLink"
			:class="{ sidebartab: true, 'sidebartab-active': linkActiveOrNot }"
		>
			<span>{{ icon }}</span>
			<h3 v-if="expanded">
				{{ title }}
			</h3>
		</router-link>
		<!-- 將按鈕移出 router-link 並在父元素中進行佈局 -->
		<button @click="toggleFavorite">
			<span class="fav" :class="{ 'fav-active': isFavorited }"
				>favorite</span
			>
		</button>

		<button v-if="title == '防戰及應變'" @click="goToURL">
			<span class="fav" :class="{ 'fav-active': isFavorited }">chat</span>
		</button>
	</div>
</template>

<style scoped lang="scss">
.sidebar-item {
	display: flex;
	align-items: center;
}
.fav {
	font-family: var(--dashboardcomponent-font-icon);
	padding: var(--font-s);
}

.fav-active {
	color: rgb(255, 65, 44);
}

.sidebartab {
	max-height: var(--font-xl);
	display: flex;
	align-items: center;
	margin: var(--font-s) 0;
	border-left: solid 4px transparent;
	border-radius: 0 5px 5px 0;
	transition: background-color 0.2s;
	white-space: nowrap;
	text-wrap: nowrap;

	&:hover {
		background-color: var(--color-component-background);
	}

	span {
		min-width: var(--font-l);
		margin-left: var(--font-s);
		font-family: var(--font-icon);
		font-size: calc(var(--font-m) * var(--font-to-icon));
	}

	h3 {
		margin-left: var(--font-s);
		font-size: var(--font-m);
		font-weight: 400;
	}

	&-active {
		border-left-color: var(--color-highlight);
		background-color: var(--color-component-background);

		span,
		h3 {
			color: var(--color-highlight);
		}
	}
}
</style>
