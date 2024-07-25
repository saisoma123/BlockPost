import React, { useState } from 'react';
import styled from 'styled-components';


// CSS code for self contained HTML code
const Container = styled.div`
  width: 100%;
  max-width: 600px;
  margin: auto;
`;

const Input = styled.input`
  width: calc(100% - 22px);
  padding: 10px;
  margin-bottom: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
`;

const Button = styled.button`
  padding: 10px 15px;
  background: #28a745;
  border: none;
  color: white;
  border-radius: 4px;
  cursor: pointer;
  &:hover {
    background: #218838;
  }
`;

const Result = styled.div`
  margin-top: 20px;
  padding: 10px;
  background: #e9ecef;
  border: 1px solid #ced4da;
  border-radius: 4px;
`;

const MessageForm = () => {
  const [creator, setCreator] = useState('');
  const [message, setMessage] = useState('');
  const [response, setResponse] = useState('');
  const [messageId, setMessageId] = useState('');
  const [retrievedMessage, setRetrievedMessage] = useState('');
  const [queryId, setQueryId] = useState('');
  const [newAccountName, setNewAccountName] = useState('');
  const [newAccountResponse, setNewAccountResponse] = useState('');

  /* This sends a POST request to the js server, which in turn stores a message
     via the BlockPost binary with a transaction that calls the Keeper addMessage
     method
  */
  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch('http://localhost:5000/store-message', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message, creator }),
      });
      const storeResponse = await res.json();
      setMessageId(storeResponse.messageId);
      setResponse('Message stored successfully!');
    } catch (error) {
      console.error('Error submitting message:', error);
      setResponse('Error submitting message.');
    }
  };

  /* This sends a request to the js server to retrieve the message from the AVL
     tree via the Querier implementation
  */

  const handleRetrieve = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch(`http://localhost:5000/get-message/${queryId}`);
      const retrieveResponse = await res.json();
      setRetrievedMessage(retrieveResponse.message);
    } catch (error) {
      console.error('Error retrieving message:', error);
    }
  };
  /* This sends a POST request to create an account on the blockchain for
     sending/storing messages, and fills the account with tokens to participate
     in on-chain transactions
  */
  const handleCreateAccount = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch('http://localhost:5000/create-account', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name: newAccountName }),
      });
      const createResponse = await res.json();
      setNewAccountResponse(`Account created successfully! Address: ${createResponse.address}`);
    } catch (error) {
      console.error('Error creating account:', error);
      setNewAccountResponse('Error creating account.');
    }
  };

  // This creates an HTML file with all the necessary fields for user interactions
  // with the BlockPost app, and calls the methods that request the js_server response
  return (
    <Container>
      <h2>Create a New Account</h2>
      <form onSubmit={handleCreateAccount}>
        <Input
          type="text"
          value={newAccountName}
          onChange={(e) => setNewAccountName(e.target.value)}
          placeholder="Enter new account name"
          required
        />
        <Button type="submit">Create Account</Button>
      </form>
      {newAccountResponse && <Result>{newAccountResponse}</Result>}

      <h2>Submit a Message</h2>
      <form onSubmit={handleSubmit}>
        <Input
          type="text"
          value={creator}
          onChange={(e) => setCreator(e.target.value)}
          placeholder="Enter your address"
          required
        />
        <Input
          type="text"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          placeholder="Enter your message"
          required
        />
        <Button type="submit">Submit</Button>
      </form>
      {response && <Result>Response: {response}</Result>}
      {messageId && <Result>Message ID: {messageId}</Result>}

      <h2>Retrieve a Message</h2>
      <form onSubmit={handleRetrieve}>
        <Input
          type="text"
          value={queryId}
          onChange={(e) => setQueryId(e.target.value)}
          placeholder="Enter message ID"
          required
        />
        <Button type="submit">Retrieve</Button>
      </form>
      {retrievedMessage && <Result>Message: {retrievedMessage}</Result>}
    </Container>
  );
};

export default MessageForm;
