const { subtle } = globalThis.crypto;

// TODO: Negalėjau hashint naudojant SHA-512, kadangi "backend" neleido
async function hashWithSHA256(value: string) {
	const textEncoder = new TextEncoder();
	const encodedValue = textEncoder.encode(value);
	const hashAsArrayBuffer = await subtle.digest('SHA-256', encodedValue);

	const hashBytes = new Uint8Array(hashAsArrayBuffer);

	return hashBytes.reduce((hex, byte) => hex + byte.toString(16).padStart(2, '0'), '');
}

export { hashWithSHA256 };
