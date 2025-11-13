import * as React from 'react'
import { Container, Box } from '@mui/material'

import style from './layout.module.scss'

interface LayoutProps {
	children: React.ReactNode
}

const Layout = ({ children }: LayoutProps) => {
	return (
		<div className={style.page}>
			<Container maxWidth="sm" className={style.wrapper}>
				<Box component="main" className={style.main}>
					{children}
				</Box>
			</Container>
		</div>
	)
}

export default Layout
