module.exports =

  actions : require './actions'
  getters : require './getters'
  stores  : [
    require './stores/teamstore'
    require './stores/teammembersidstore'
    require './stores/teammembersrolestore'
    require './stores/teaminvitationinputvaluesstore'
    require './stores/teaminvitationstore'
    require './stores/teamsearchinputvaluestore'
    require './stores/teamdisabledmembersstore'
    require './stores/teamallmembersloadedstore'
    require './stores/teamsendinvitesfocusstore'
  ]

  register: (reactor) -> reactor.registerStores @stores
