const AWS = require('aws-sdk');
var log4js = require('log4js');
var logger = log4js.getLogger();
logger.level = 'debug';


AWS.config.update({ region: 'ap-northeast-1' });
const dynamodb = new AWS.DynamoDB.DocumentClient();

const TableName = 'performance-test-dynamodb';

(async () => {
  try {
    logger.debug(TableName);

    const params = {
      TableName,
      Key: {
        id: '10'
      }
    };

    const documentClient = new AWS.DynamoDB.DocumentClient();

    const start = new Date().getTime(); 
    const result = await documentClient.get(params).promise();
    const end = new Date().getTime(); 
    logger.debug(result);
    logger.info(`${end-start},ms`);

  } catch (error) {
    logger.error(error);
  }
})()
