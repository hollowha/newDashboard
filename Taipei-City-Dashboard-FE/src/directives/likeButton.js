// export default {
// 	mounted(el, binding) {
// 		const likeButton = document.createElement("button");
// 		likeButton.innerText = "Like";
// 		likeButton.className = "like-button";
// 		likeButton.style.position = "absolute";
// 		likeButton.style.right = "10px";
// 		likeButton.style.bottom = "10px";

// 		likeButton.addEventListener("click", () => {
// 			alert(`Liked component with id: ${binding.value}`);
// 		});

// 		el.style.position = "relative"; // 確保父元素可以相對定位子元素
// 		el.appendChild(likeButton);
// 	},
// };
// directives/likeButton.js
export default {
	mounted(el, binding) {
		const likeButton = document.createElement("button");
		likeButton.innerText = "Like";
		likeButton.className = "like-button";

		likeButton.addEventListener("click", () => {
			alert(`Liked component with id: ${binding.value}`);
		});

		// 確保包裹的 div 元素不影響內部組件的佈局
		el.style.position = "relative";
		el.style.display = "inline-block";

		el.appendChild(likeButton);
	},
};
