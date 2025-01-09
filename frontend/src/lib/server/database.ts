import { VITE_SERVER_ADDRESS } from '$env/static/private';
import type { Coffee } from '../coffee';

async function getCoffees(sessionToken: string) {
	const response = await fetch(`${VITE_SERVER_ADDRESS}/coffeeinventory`, {
		method: 'GET',
		headers: {
			Cookie: `session_token=${sessionToken}`
		}
	});

	const { coffeeList }: { coffeeList: Coffee[] } = await response.json();

	return coffeeList ?? [];
}

async function getCoffee(sessionToken: string, id: number) {
	const response = await fetch(`${VITE_SERVER_ADDRESS}/fetchcoffee/${id}`, {
		method: 'POST',
		headers: {
			Cookie: `session_token=${sessionToken}`
		}
	});

	const { coffeeObject }: { coffeeObject: Coffee } = await response.json();

	return coffeeObject ?? [];
}

async function getOrders(sessionToken: string) {
	const response = await fetch(`${VITE_SERVER_ADDRESS}/orderlist`, {
		method: 'GET',
		headers: {
			Cookie: `session_token=${sessionToken}`
		}
	});

	// const { coffeeList, title }: { coffeeList: Coffee[]; title: string } = await response.json();

	return await response.json();
}

async function getAllOrders(sessionToken: string) {
	const response = await fetch(`${VITE_SERVER_ADDRESS}/allorderlist`, {
		method: 'GET',
		headers: {
			Cookie: `session_token=${sessionToken}`
		}
	});

	// const { coffeeList, title }: { coffeeList: Coffee[]; title: string } = await response.json();

	return await response.json();
}

// TODO
async function getUsers(sessionToken: string) {
	const response = await fetch(`${VITE_SERVER_ADDRESS}/fetchallusers/`, {
		method: 'GET',
		headers: {
			Cookie: `session_token=${sessionToken}`
		}
	});

	const { users } = await response.json();

	return users;
}

async function getUserBySessionToken(sessionToken: string) {
	// protected.POST("/fetchuser/:session_token", endpoints.GetSingleUser(db))

	const response = await fetch(`${VITE_SERVER_ADDRESS}/fetchuser/${sessionToken}`, {
		method: 'POST',
		headers: {
			Cookie: `session_token=${sessionToken}`
		}
	});

	const { userObject } = await response.json();

	return userObject;
}

async function getUserOrders(sessionToken: string, id: number) {
	const response = await fetch(`${VITE_SERVER_ADDRESS}/fetchuserorder/${id}`, {
		method: 'POST',
		headers: {
			Cookie: `session_token=${sessionToken}`
		}
	});

	const { orders } = await response.json();

	return orders ?? [];
}

async function getOrderItems(sessionToken: string, id: number) {
	const response = await fetch(`${VITE_SERVER_ADDRESS}/getorderitems/${id}`, {
		method: 'POST',
		headers: {
			Cookie: `session_token=${sessionToken}`
		}
	});

	const body = await response.json();

	return body ?? [];
}

export {
	getCoffees,
	getCoffee,
	getOrders,
	getAllOrders,
	getUsers,
	getUserBySessionToken,
	getUserOrders,
	getOrderItems
};
