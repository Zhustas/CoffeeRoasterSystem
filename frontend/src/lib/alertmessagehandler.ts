import { writable } from 'svelte/store';

interface AlertMessageHandler {
	show: boolean;
	status: string;
	message: string;
}

const initialState: AlertMessageHandler = { show: false, status: '', message: '' };
const alertStore = writable<AlertMessageHandler>(initialState);

function showAlertMessage(status: string, message: string) {
	alertStore.update((state) => {
		if (state.show) {
			return state;
		}

		return { show: true, status, message };
	});

	setTimeout(() => {
		alertStore.update(() => ({ ...initialState }));
	}, 5000);
}

export { type AlertMessageHandler, alertStore, showAlertMessage };
