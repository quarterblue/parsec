import React from 'react';
import { Input, Card, Button } from 'antd';

const { TextArea } = Input;

const Block = () => {
  return (
    <Card
      title="Default size card"
      extra={<a href="#">More</a>}
      style={{ width: 500, margin: '3em' }}
    >
      <p>Block:</p>
      <Input />
      <br />
      <p>Nonce:</p>
      <Input />
      <p>Data:</p>
      <TextArea rows={3} />
      <p>Previous Hash:</p>
      <Input />
      <p>Hash:</p>
      <Input />
      <Button style={{ marginTop: '1em' }} type="primary">
        Mine
      </Button>
    </Card>
  );
};

export default Block;
