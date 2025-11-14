import { Provider } from 'react-redux'
import { store } from './store/store'
import Layout from './component/layout'
import Theme from './presentation/theme'
import Router from './infra/Router'
import { QueryClientProvider } from '@tanstack/react-query'
import { queryClient } from 'utils/api'

function App() {
	return (
		<Provider store={store}>
			<Theme>
				<Layout>
					<QueryClientProvider client={queryClient}>
						<Router />
					</QueryClientProvider>
				</Layout>
			</Theme>
		</Provider>
	)
}

export default App
