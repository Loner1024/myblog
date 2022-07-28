import type { NextPage } from 'next'
import Header from '../components/header'
import HomeProfile from '../components/homeProfile'
import { Page, Text, Card } from '@geist-ui/core'
import { Car } from '@icon-park/react'

const Home: NextPage = () => {
  return (
    <Page className='w-full' width='100%'>
      <Page.Header width='100%'>
        <Header />
      </Page.Header>
      <Page.Content>
        <HomeProfile />
      </Page.Content>
      <Page.Footer>
        <Text h1>fonter</Text>
      </Page.Footer>
    </Page>
  )
}

export default Home
