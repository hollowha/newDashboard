<script setup>
import { ref, computed } from "vue";
import { useDialogStore } from "../../store/dialogStore";

import { jsonToCsv } from "../../assets/utilityFunctions/jsonToCsv";
import DialogContainer from "./DialogContainer.vue";

const dialogStore = useDialogStore();

// Stores the inputted dashboard name
const name = ref(dialogStore.moreInfoContent.name);
// Stores the file type
const fileType = ref("JSON");

const parsedJson = computed(() => {
	let json = {};
	json.data = dialogStore.moreInfoContent.chart_data;
	if (dialogStore.moreInfoContent.chart_config.categories) {
		json.categories = dialogStore.moreInfoContent.chart_config.categories;
	}

	const jsonString = encodeURIComponent(JSON.stringify(json));
	// const base64Json = btoa(jsonString)
	return jsonString;
});

const parsedCsv = computed(() => {
	const csvString = dialogStore.moreInfoContent.chart_data
		? jsonToCsv(
				dialogStore.moreInfoContent.chart_data,
				dialogStore.moreInfoContent.chart_config
		  )
		: "";
	return encodeURI(csvString);
});

function handleSubmit() {
	handleClose();
}
function handleClose() {
	name.value = dialogStore.moreInfoContent.name;
	dialogStore.dialogs.downloadData = false;
}

const dataContent = ref("");

// onMounted(() => {
// 	fetchData();
// });

// function fetchData() {
// 	// const url = `data:application/json;charset=utf-8,%7B%22data%22%3A%5B%7B%22name%22%3A%22%E7%99%BC%E7%94%9F%E4%BB%B6%E6%95%B8%22%2C%22data%22%3A%5B%7B%22x%22%3A%222023-06-01T00%3A00%3A00%2B08%3A00%22%2C%22y%22%3A3685%7D%2C%7B%22x%22%3A%222023-07-01T00%3A00%3A00%2B08%3A00%22%2C%22y%22%3A4623%7D%2C%7B%22x%22%3A%222023-08-01T00%3A00%3A00%2B08%3A00%22%2C%22y%22%3A4534%7D%5D%7D%2C%7B%22name%22%3A%22%E7%A0%B4%E7%8D%B2%E4%BB%B6%E6%95%B8%22%2C%22data%22%3A%5B%7B%22x%22%3A%222023-06-01T00%3A00%3A00%2B08%3A00%22%2C%22y%22%3A2793%7D%2C%7B%22x%22%3A%222023-07-01T00%3A00%3A00%2B08%3A00%22%2C%22y%22%3A3701%7D%2C%7B%22x%22%3A%222023-08-01T00%3A00%3A00%2B08%3A00%22%2C%22y%22%3A4109%7D%5D%7D%5D%7D`;
// 	const url = `data:application/json;charset=utf-8,%7B%22data%22%3A%5B%7B%22name%22%3A%22%E5%95%9F%E5%8B%95%E4%B8%AD%22%2C%22icon%22%3A%22%22%2C%22data%22%3A%5B2%5D%7D%2C%7B%22name%22%3A%22%E6%9C%AA%E5%95%9F%E5%8B%95%22%2C%22icon%22%3A%22%22%2C%22data%22%3A%5B72%5D%7D%5D%2C%22categories%22%3A%5B%22%E5%95%9F%E5%8B%95%E6%8A%BD%E6%B0%B4%E7%AB%99%22%5D%7D`;
// 	fetch(url)
// 		.then((response) => response.json())
// 		.then((data) => {
// 			dataContent.value = JSON.stringify(data, null, 2);
// 		});
// }

try {
	// Decode the URI to get the JSON string
	const dataUri = parsedJson.value;
	const base64EncodedData = dataUri.split(",")[1]; // Assuming the data is base64 encoded after 'data:application/json;base64,'
	const jsonStr = decodeURIComponent(atob(base64EncodedData)); // Decoding base64 and then URI component
	dataContent.value = JSON.stringify(JSON.parse(jsonStr), null, 2);
} catch (error) {
	dataContent.value = "Error parsing JSON from URI: " + error.message;
}

const displayContent = computed(() => {
	if (fileType.value === "JSON") {
		try {
			const decodedJson = decodeURIComponent(parsedJson.value);
			return JSON.stringify(JSON.parse(decodedJson), null, 2);
		} catch (e) {
			return `Error decoding JSON data: ${e.message}`;
		}
	} else if (fileType.value === "CSV") {
		try {
			// Assuming CSV data needs no further decoding
			return decodeURIComponent(parsedCsv.value);
		} catch (e) {
			return `Error processing CSV data: ${e.message}`;
		}
	}
	return "";
});
</script>

<template>
	<DialogContainer :dialog="`downloadData`" @on-close="handleClose">
		<div class="downloaddata">
			<h2>下載資料</h2>
			<div class="downloaddata-input">
				<h3>請輸入檔名</h3>
				<input v-model="name" type="text" :minlength="1" required />
			</div>
			<h3>請選擇檔案格式</h3>
			<!-- 在這邊預覽data -->
			<div style="height: 200px; width: 100%">
				<div class="data-preview">
					<pre>{{ displayContent }}</pre>
				</div>
			</div>
			<div>
				<input
					id="JSON"
					v-model="fileType"
					class="downloaddata-radio"
					type="radio"
					value="JSON"
				/>
				<label for="JSON">
					<div />
					JSON
				</label>
				<input
					id="CSV"
					v-model="fileType"
					class="downloaddata-radio"
					type="radio"
					value="CSV"
				/>
				<label for="CSV">
					<div />
					CSV (UTF-8)
				</label>
			</div>
			<div class="downloaddata-control">
				<button
					class="downloaddata-control-cancel"
					@click="handleClose"
				>
					取消
				</button>
				<button
					v-if="name && fileType === 'JSON'"
					class="downloaddata-control-confirm"
					@click="handleSubmit"
				>
					<a
						:href="`data:application/json;charset=utf-8,${parsedJson}`"
						:download="`${name}.json`"
						>下載JSON</a
					>
				</button>
				<button
					v-if="name && fileType === 'CSV'"
					class="downloaddata-control-confirm"
					@click="handleSubmit"
				>
					<a
						:href="`data:text/csv;charset=utf-8,${parsedCsv}`"
						:download="`${name}.csv`"
						>下載CSV</a
					>
				</button>
			</div>
		</div>
	</DialogContainer>
</template>

<style scoped lang="scss">
.downloaddata {
	width: 300px;

	h3 {
		margin-bottom: 0.5rem;
		font-size: var(--font-s);
		font-weight: 400;
		color: var(--color-complement-text);
	}

	&-input {
		display: flex;
		flex-direction: column;
		margin: 0.5rem 0;
	}

	&-radio {
		display: none;

		&:checked + label {
			color: white;

			div {
				background-color: var(--color-highlight);
			}
		}

		&:hover + label {
			color: var(--color-highlight);

			div {
				border-color: var(--color-highlight);
			}
		}
	}

	label {
		position: relative;
		display: flex;
		align-items: center;
		margin-bottom: 2px;
		font-size: var(--font-s);
		color: var(--color-complement-text);
		transition: color 0.2s;
		cursor: pointer;

		div {
			width: calc(var(--font-s) / 2);
			height: calc(var(--font-s) / 2);
			margin-right: 4px;
			padding: calc(var(--font-s) / 4);
			border-radius: 50%;
			border: 1px solid var(--color-border);
			transition: background-color 0.2s;
		}
	}

	&-control {
		display: flex;
		justify-content: flex-end;

		&-cancel {
			margin: 0 2px;
			padding: 4px 6px;
			border-radius: 5px;
			transition: color 0.2s;

			&:hover {
				color: var(--color-highlight);
			}
		}

		&-confirm {
			margin: 0 2px;
			padding: 4px 10px;
			border-radius: 5px;
			background-color: var(--color-highlight);
			transition: opacity 0.2s;

			&:hover {
				opacity: 0.8;
			}
		}
	}

	.data-preview {
		height: 150px; /* Adjusted for demonstration */
		width: 100%; /* Adjusted to take full width of its container */
		overflow: auto; /* Enables scrolling for both axes if content overflows */
		background-color: var(--background-color); /* Optional: for better UI */
		/* padding: 10px 30px; Optional: for spacing */
		margin: 0px 20px;
		white-space: pre; /* Keeps text formatting and enables horizontal scrolling */
		/* border: 3px solid var(--color-highlight); Optional: for better UI */
	}
}
</style>
