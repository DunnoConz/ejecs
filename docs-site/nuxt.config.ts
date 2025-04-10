// https://nuxt.com/docs/api/configuration/nuxt-config
import { defineNuxtConfig } from 'nuxt/config'

export default defineNuxtConfig({
  devtools: { enabled: true },

  modules: [
    '@nuxt/content',
    '@nuxtjs/tailwindcss',
  ],

  app: {
    head: {
      title: 'EJECS Documentation',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Documentation for the EJECS Entity Component System' }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
      ]
    }
  },

  // @ts-ignore
  content: {
    highlight: {
      theme: 'github-light',
      preload: ['cpp', 'typescript', 'bash']
    },
    markdown: {
      toc: {
        depth: 4,
        searchDepth: 4
      }
    }
  },

  ssr: true,
  nitro: {
    preset: 'cloudflare-pages'
  },

  compatibilityDate: '2025-04-09'
})