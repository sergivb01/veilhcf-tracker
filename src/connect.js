const config = require('../config.json');

function connectToServer(bot) {
  setTimeout(() => {
    bot.setQuickBarSlot(4);
    bot.activateItem();

    setTimeout(() => {
      bot.clickWindow(config.server.slot, 0, 0, err => {
        if (err) console.error(err);

        console.log(config.server.slot);

        console.info('A Menu has been clicked');
      });
    }, 1500);
  }, 1500);
}

module.exports = connectToServer;
