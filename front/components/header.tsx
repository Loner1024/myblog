import React, { Component } from 'react'
import { Page, Tabs, Grid, Card } from '@geist-ui/core'
import { GrinningFaceWithTightlyClosedEyes } from '@icon-park/react'

interface Menu {
  id: number
  label: string
  url: string
}

export default class Header extends Component {
  render() {
    const menus: Menu[] = [
      {
        id: 1,
        label: 'Home',
        url: 'home'
      },
      {
        id: 2,
        label: 'About',
        url: 'about'
      },
      {
        id: 3,
        label: 'Link',
        url: 'link'
      }
    ]
    const menuComponent = menus.map(menu => (
      <Tabs.Item
        font='18px'
        key={menu.id}
        value={menu.url}
        label={menu.label}
        className='m-0'
      ></Tabs.Item>
    ))
    return (
      <div className='mt-5 w-auto h-1/6'>
        <Grid.Container gap={0} alignContent='center' alignItems='center'>
          <Grid
            xs={12}
            justify='flex-start'
            alignContent='center'
            alignItems='center'
          >
            <GrinningFaceWithTightlyClosedEyes
              theme='outline'
              fill='#333'
              className='text-5xl'
            />
          </Grid>
          <Grid xs={6} alignContent='center' alignItems='center'>
            <Tabs align='center' hideDivider hideBorder>
              {menuComponent}
            </Tabs>
          </Grid>
          <Grid xs></Grid>
        </Grid.Container>
      </div>
    )
  }
}
