// https://nuxt.com/docs/api/configuration/nuxt-config
import { defineNuxtConfig } from 'nuxt/config'

export default defineNuxtConfig({
  devtools: { enabled: true },

  modules: [
    '@nuxt/content',
    '@nuxtjs/tailwindcss',
  ],

  app: {
    baseURL: '/ejecs/', // GitHub Pages repository name
    head: {
      title: 'EJECS Documentation',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Documentation for the EJECS Entity Component System for Roblox' }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
      ]
    }
  },

  content: {
    highlight: {
      theme: 'github-light',
      preload: ['lua', 'typescript']
    },
    markdown: {
      toc: {
        depth: 4,
        searchDepth: 4
      }
    }
  },

  compatibilityDate: '2025-04-09'
})