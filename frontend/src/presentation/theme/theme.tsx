import * as React from 'react'
import { createTheme, ThemeProvider } from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'
import colors from './colors.module.scss'

const theme = createTheme({
	cssVariables: true,
	colorSchemes: {
		light: {
			palette: {
				primary: {
					light: colors.orange300,
					main: colors.orange500,
					dark: colors.orange700,
					contrastText: '#fff',
				},
				secondary: {
					light: colors.blue300,
					main: colors.blue500,
					dark: colors.blue700,
					contrastText: '#fff',
				},
				background: {
					default: '#ffffff',
					paper: colors.spaceGrey100,
				},
				text: {
					primary: '#1a1a1a',
					secondary: colors.spaceGrey500,
				},
			},
		},
		dark: {
			palette: {
				primary: {
					light: colors.orange400,
					main: colors.orange500,
					dark: colors.orange800,
					contrastText: '#fff',
				},
				secondary: {
					light: colors.blue400,
					main: colors.blue500,
					dark: colors.blue800,
					contrastText: '#fff',
				},
				background: {
					default: '#121212',
					paper: '#1e1e1e',
				},
				text: {
					primary: '#ffffff',
					secondary: colors.spaceGrey400,
				},
			},
		},
	},
})

interface ThemeProps {
	children: React.ReactNode
}

export const Theme = ({ children }: ThemeProps) => (
	<ThemeProvider theme={theme}>
		<CssBaseline />
		{children}
	</ThemeProvider>
)

export default Theme