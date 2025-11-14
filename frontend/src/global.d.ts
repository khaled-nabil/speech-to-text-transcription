/// <reference types="vitest/globals" />

declare module '*.scss' {
	interface ScssColors {
		orange50: string
		orange100: string
		orange200: string
		orange300: string
		orange400: string
		orange500: string
		orange600: string
		orange700: string
		orange800: string
		orange900: string
		spaceGrey100: string
		spaceGrey200: string
		spaceGrey300: string
		spaceGrey400: string
		spaceGrey500: string
	}

	type ClassMap = { [className: string]: string }

	const styles: ClassMap & ScssColors

	export const colors: ScssColors

	export default styles
}
