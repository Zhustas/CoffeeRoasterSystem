import { writable } from 'svelte/store';

export const wantsToAddCoffee = writable(false);
export const wantsToCancelOrder = writable(false);
export const wantsToMakeOrder = writable(false);

export const wantsToAddUser = writable(false);
export const userAddedStatus = writable(0);
