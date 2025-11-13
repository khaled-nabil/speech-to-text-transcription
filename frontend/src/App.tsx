import { Provider } from 'react-redux'
import { store } from './store/store'
import Layout from './component/layout'
import Theme from './presentation/theme'
import Router from './infra/Router'

function App() {
	return (
		<Provider store={store}>
			<Theme>
				<Layout>
					<Router />
				</Layout>
			</Theme>
		</Provider>
	)
}

export default App
