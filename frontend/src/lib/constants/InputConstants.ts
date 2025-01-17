export const VALID_PASSWORD_LENGTH = 8;
export const VALID_MIN_USERNAME_LENGTH = 5;
export const VALID_MAX_USERNAME_LENGTH = 30;
export const VALID_MAX_NAME_LENGTH = 50;
export const VALID_MAX_LASTNAME_LENGTH = 50;

export const ERROR_EMPTY_INPUT = 'Užpildykite lauką.';
export const ERROR_PASSWORDS_NOT_MATCH = 'Slaptažodžiai nesutampa.';
export const ERROR_PASSWORD_TOO_SHORT = `Slaptažodis turi būti sudarytas bent iš ${VALID_PASSWORD_LENGTH} simbolių.`;
export const ERROR_USERNAME_BAD_LENGTH = `Vartotojo vardas turi būti nuo ${VALID_MIN_USERNAME_LENGTH} iki ${VALID_MAX_USERNAME_LENGTH} simbolių.`;
export const ERROR_EMAIL_INVALID = 'Blogas elektroninio pašto formatas.';
export const ERROR_NAME_BAD_LENGTH = `Vardas turi būti iki ${VALID_MAX_NAME_LENGTH} simbolių.`;
export const ERROR_LASTNAME_BAD_LENGTH = `Lauko ilgis turi būti iki ${VALID_MAX_LASTNAME_LENGTH} simbolių.`;
export const ERROR_CONSISTS_NOT_ONLY_LETTERS = 'Laukas turi būti sudarytas tik iš raidžių.';
export const ERROR_CONSISTS_NOT_ONLY_LETTERS_NUMBERS =
	'Laukas turi būti sudarytas tik iš raidžių ir skaičių.';
