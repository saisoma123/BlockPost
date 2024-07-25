const express = require('express');
const bodyParser = require('body-parser');
const axios = require('axios');
const { exec } = require('child_process');
const cors = require('cors');

const app = express();
const port = 5000;

app.use(cors());
app.use(bodyParser.json());

/* This function queries the transaction and gets the message id from the event 
   logs, it does it multiple times, as the transaction can take a while to be 
   fully registered on-chain
*/

const delay = ms => new Promise(resolve => setTimeout(resolve, ms));

const queryTransaction = async (txHash, retries = 10, delayTime = 5000) => {
  for (let i = 0; i < retries; i++) {
    try {
      const { stdout, stderr } = await new Promise((resolve, reject) => {
        exec(`BlockPostd query tx ${txHash} -o json | jq -r '.events[] | select(.type=="message").attributes[] | select(.key=="message_id").value'`, (error, stdout, stderr) => {
          if (error) {
            if (stderr.includes("not found")) {
              console.log(`Attempt ${i + 1}: Transaction not found, retrying...`);
              resolve({ stdout: null, stderr });
            } else {
              reject({ error, stderr });
            }
          } else {
            resolve({ stdout, stderr });
          }
        });
      });

      if (stdout) {
        console.log(`Attempt ${i + 1}: Transaction found`);
        return stdout.trim();
      } else {
        await delay(delayTime);
      }
    } catch (err) {
      console.log(`Attempt ${i + 1} failed: ${err.stderr || err.error}`);
      if (i < retries - 1) {
        await delay(delayTime);
      } else {
        throw new Error(`Failed to query transaction after ${retries} attempts`);
      }
    }
  }
};

/* This directly runs the BlockPost binary app and calls the Keeper add message
   function, which stores the message in an IAVL tree,
  and calls the queryTransaction method to retrieve the message id
   from the event logs
*/
app.post('/store-message', (req, res) => {
  console.log('Received a request to /store-message');
  const { message, creator } = req.body;

  exec(`BlockPostd tx blockpost block-post-message "${creator}" "${message}" --chain-id BlockPost --yes`, async (error, stdout, stderr) => {
    if (error) {
      console.error(`Error storing message: ${stderr}`);
      return res.status(500).json({ error: stderr });
    }

    const txHashMatch = stdout.match(/txhash: (.+)/);
    if (!txHashMatch) {
      console.error('Transaction hash not found in stdout');
      return res.status(500).json({ error: 'Transaction hash not found' });
    }

    const txHash = txHashMatch[1].trim();
    console.log(`Message stored, transaction hash: ${txHash}`);

    try {
      const messageId = await queryTransaction(txHash);
      console.log(`Message ID: ${messageId}`);
      res.json({ messageId });
    } catch (err) {
      console.error(`Error waiting for transaction: ${err.message}`);
      res.status(500).json({ error: err.message });
    }
  });
});

// This creates an account on the BlockPost chain and sends tokens for the 
// account, so that they can participate in on-chain transactions
app.post('/create-account', (req, res) => {
  console.log('Received a request to /create-account');
  const { name } = req.body;

  // Creates a new account and outputs an address to the user
  exec(`BlockPostd keys add "${name}"`, (error, stdout, stderr) => {
    if (error) {
      console.error(`Error creating account: ${stderr}`);
      return res.status(500).json({ error: stderr });
    }
    console.log(stdout);
    const address = stdout.match(/address: (.+)/)[1];
    console.log(`Account created, address: ${address}`);

    // Send tokens to the newly created account
    exec(`BlockPostd tx bank send cosmos1uu7asml6ktx8kpn8akt4y237kty9jmqwea9wea ${address} 1token --chain-id BlockPost --yes`, (error, stdout, stderr) => {
      if (error) {
        console.error(`Error sending tokens: ${stderr}`);
        return res.status(500).json({ error: stderr });
      }
      console.log(`Tokens sent to ${address}: ${stdout}`);
      res.json({ address });
    });
  });
});

// This calls the Querier function and retrieves the message from the IAVL tree
app.get('/get-message/:id', async (req, res) => {
  console.log(`Received a request to /get-message/${req.params.id}`);
  const { id } = req.params;

  try {
    const response = await axios.get(`http://localhost:1317/saisoma123/BlockPost/blockpost/messages/${id}`);
    const message = response.data.message;
    console.log(`Message retrieved: ${message}`);
    res.json({ message });
  } catch (error) {
    console.error(`Error retrieving message: ${error.message}`);
    res.status(500).json({ error: error.message });
  }
});

app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});
