import { Provider } from 'react-redux'
import { store } from './store/store'
import Theme from './presentation/theme'
import Router from './infra/Router'
import { QueryClientProvider } from '@tanstack/react-query'
import { queryClient } from 'utils/api'

function App() {
	return (
		<Provider store={store}>
			<Theme>
				<QueryClientProvider client={queryClient}>
					<Router />
				</QueryClientProvider>
			</Theme>
		</Provider>
	)
}

export default App
