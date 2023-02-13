module.exports = {
  theme: 'DeltaChain',
  title: 'DeltaChain Core',
  base: process.env.VUEPRESS_BASE,
  themeConfig: {
    repo: 'DeltaChain/DeltaChain',
    docsRepo: 'DeltaChain/DeltaChain',
    docsDir: 'docs',
    editLinks: true,
    label: 'core',
    algolia: {
      id: "BH4D9OD16A",
      key: "59f0e2deb984aa9cdf2b3a5fd24ac501",
      index: "DeltaChain"
    },
    versions: [
      {
        "label": "v0.34 (latest)",
        "key": "v0.34"
      },
      {
        "label": "v0.33",
        "key": "v0.33"
      }
    ],
    topbar: {
      banner: false,
    },
    sidebar: {
      auto: true,
      nav: [
        {
          title: 'Resources',
          children: [
            {
              title: 'RPC',
              path: (process.env.VUEPRESS_BASE ? process.env.VUEPRESS_BASE : '/')+'rpc/',
              static: true
            },
          ]
        }
      ]
    },
    gutter: {
      title: 'Help & Support',
      editLink: true,
      forum: {
        title: 'DeltaChain Discussions',
        text: 'Join the DeltaChain discussions to learn more',
        url: 'https://github.com/DeltaChain/DeltaChain/discussions',
        bg: '#0B7E0B',
        logo: 'DeltaChain'
      },
      github: {
        title: 'Found an Issue?',
        text: 'Help us improve this page by suggesting edits on GitHub.'
      }
    },
    footer: {
      question: {
        text: 'Chat with DeltaChain developers in <a href=\'https://discord.gg/vcExX9T\' target=\'_blank\'>Discord</a> or reach out on <a href=\'https://github.com/DeltaChain/DeltaChain/discussions\' target=\'_blank\'>GitHub</a> to learn more.'
      },
      logo: '/logo-bw.svg',
      textLink: {
        text: 'DeltaChain.com',
        url: 'https://DeltaChain.com'
      },
      services: [
        {
          service: 'medium',
          url: 'https://medium.com/@DeltaChain'
        },
        {
          service: 'twitter',
          url: 'https://twitter.com/DeltaChain_team'
        },
        {
          service: 'linkedin',
          url: 'https://www.linkedin.com/company/DeltaChain/'
        },
        {
          service: 'reddit',
          url: 'https://reddit.com/r/DeltaChainnetwork'
        },
        {
          service: 'telegram',
          url: 'https://t.me/DeltaChainproject'
        },
        {
          service: 'youtube',
          url: 'https://www.youtube.com/c/DeltaChainProject'
        }
      ],
      smallprint:
        'The development of DeltaChain Core is led primarily by [Interchain GmbH](https://interchain.berlin/). Funding for this development comes primarily from the Interchain Foundation, a Swiss non-profit. The DeltaChain trademark is owned by DeltaChain Inc, the for-profit entity that also maintains this website.',
      links: [
        {
          title: 'Documentation',
          children: [
            {
              title: 'DeltaChain SDK',
              url: 'https://docs.DeltaChain.network'
            },
            {
              title: 'DeltaChain Hub',
              url: 'https://hub.DeltaChain.network'
            }
          ]
        },
        {
          title: 'Community',
          children: [
            {
              title: 'DeltaChain blog',
              url: 'https://medium.com/@DeltaChain'
            },
            {
              title: 'GitHub Discussions',
              url: 'https://github.com/DeltaChain/DeltaChain/discussions'
            }
          ]
        },
        {
          title: 'Contributing',
          children: [
            {
              title: 'Contributing to the docs',
              url: 'https://github.com/DeltaChain/DeltaChain'
            },
            {
              title: 'Source code on GitHub',
              url: 'https://github.com/DeltaChain/DeltaChain'
            },
            {
              title: 'Careers at DeltaChain',
              url: 'https://DeltaChain.com/careers'
            }
          ]
        }
      ]
    }
  },
  plugins: [
    [
      '@vuepress/google-analytics',
      {
        ga: 'UA-51029217-11'
      }
    ],
    [
      '@vuepress/plugin-html-redirect',
      {
        countdown: 0
      }
    ]
  ]
};
