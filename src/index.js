const mineflayer = require('mineflayer');
const tokens = require('prismarine-tokens');
const redis = require('redis');
const stripAnsi = require('strip-ansi');

const config = require('../config.json');
const connectToServer = require('./connect');

const client = redis.createClient(config.redis);

function createBot(options) {
  tokens.use(options, function(_err, _opts) {
    if (_err) throw _err;
    const bot = mineflayer.createBot(_opts);

    bot.on('error', err => {
      if (err.code === undefined) {
        console.log('Invalid credentials OR bot needs to wait because it relogged too quickly.');
      }
      process.exit();
    });

    bot.on('message', json => {
      const plainText = stripAnsi(json.toAnsi());
      console.info(plainText);
      client.publish('mc', plainText);
    });

    bot.once('login', () => {
      console.log('logged in');
      connectToServer(bot);
    });

    bot.on('end', function() {
      console.log('Bot has ended');
      process.exit();
    });
  });
}

createBot({
  host: config.server.host,
  port: config.server.port,
  username: config.auth.username,
  password: config.auth.password,
  version: config.server.version,
  verbose: true,
  tokensLocation: './bot_tokens.json',
  tokensDebug: true,
});
