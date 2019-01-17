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


  const documentClient = new AWS.DynamoDB.DocumentClient();

  let result;
  let sum = 0;
  for(let i=0; i<16; i=i+1){

    const start = new Date().getTime(); 
    result = await documentClient.get({
        TableName,
        Key: {
          id: `${i}`
        }
      }).promise();
    const end = new Date().getTime(); 
    logger.info(`${end-start},ms`)
    sum = sum + (end-start);
  }
  logger.debug(result.Item.id);
  logger.info(`average,${sum/16},ms`);

  } catch (error) {
  logger.error(error);
  }
})()
