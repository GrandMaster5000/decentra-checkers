// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import Tolik22869CheckersTolik22869CheckersCheckers from './tolik22869/checkers/tolik22869.checkers.checkers'


export default { 
  Tolik22869CheckersTolik22869CheckersCheckers: load(Tolik22869CheckersTolik22869CheckersCheckers, 'tolik22869.checkers.checkers'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}
