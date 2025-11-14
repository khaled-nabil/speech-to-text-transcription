import { render, screen } from '@testing-library/react'
import TranscriptionRow from '../transcriptionRow'
import type { TranscriptionResponse } from 'types/transcription'
import { v5 as uuidv5 } from 'uuid'

const transactionMock: TranscriptionResponse = {
	id: uuidv5('test-file', uuidv5.DNS),
	uploadDate: '2024-06-01T12:00:00Z',
	transcriptText: 'This is a test transcription text.',
	status: 'SUCCESS',
}

describe('TranscriptionRow', () => {
	test('renders the transcription text', () => {
		render(<TranscriptionRow transcription={transactionMock} />)
		expect(
			screen.getByText('This is a test transcription text.')
		).toBeTruthy()
	})

	test('renders the formatted date', () => {
		render(<TranscriptionRow transcription={transactionMock} />)
		const formattedDate = new Date('2024-06-01T12:00:00Z').toLocaleString()
		expect(screen.getByText(formattedDate)).toBeTruthy()
	})

	test('renders within a ListItem', () => {
		const { container } = render(
			<TranscriptionRow transcription={transactionMock} />
		)
		expect(container.querySelector('.MuiListItem-root')).toBeTruthy()
	})
})
