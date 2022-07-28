import type { AppProps } from 'next/app'
import { GeistProvider, CssBaseline } from '@geist-ui/core'
import '../styles/globals.css'

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <GeistProvider themeType='dark'>
      <CssBaseline />
      <Component {...pageProps} />
    </GeistProvider>
  )
}

export default MyApp
