const AWS = require('aws-sdk');
const data = require('./data');
var log4js = require('log4js');
var logger = log4js.getLogger();
logger.level = 'debug';


AWS.config.update({ region: 'ap-northeast-1' });
const dynamodb = new AWS.DynamoDB.DocumentClient();

const TableName = 'performance-test-dynamodb';



(async () => {
  try {
    logger.debug(TableName);

    const params = { RequestItems: { [TableName]: [] } }
    data.items.forEach(function (v, index) {
      params['RequestItems'][TableName].push({
        PutRequest: {
          Item: v
        }
      });
    })

    // logger.debug(params);

    const result = await dynamodb.batchWrite(params).promise();
    logger.debug(result);


  } catch (error) {
    logger.error(error);
  }
})()