<template>
	<div class="share-button" @click="shareUrl">
		<span>Share</span>
	</div>
</template>

<script setup>
// import { ref } from "vue";

const shareUrl = () => {
	const url = window.location.href;
	const text = "Check out this awesome page!";
	const shareData = {
		title: "Web Page",
		text: text,
		url: url,
	};

	if (navigator.share) {
		navigator
			.share(shareData)
			.then(() => {})
			.catch((error) => {
				console.error("Something went wrong sharing the URL", error);
			});
	} else {
		const fallbackUrl = `https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(
			url
		)}`;
		window.open(fallbackUrl, "_blank");
	}
};
</script>

<style scoped>
.share-button {
	position: fixed;
	bottom: 20px;
	left: 20px;
	background-color: #515151;
	color: white;
	padding: 10px 20px;
	border-radius: 50px;
	cursor: pointer;
	box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
	display: flex;
	align-items: center;
	justify-content: center;
	font-size: 16px;
	transition: background-color 0.3s;
	z-index: 1000;
}

.share-button:hover {
	background-color: #515151;
}
</style>
