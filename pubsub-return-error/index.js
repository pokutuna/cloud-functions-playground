exports.app = (message, context) => {
  const eventAge = Date.now() - Date.parse(context.timestamp);
  const eventMaxAge = 10000;
  if (eventAge > eventMaxAge) {
    console.log(`Dropping event ${context.eventId} with age ${eventAge} ms`);
    return;
  }

  const data = JSON.parse(Buffer.from(message.data, "base64").toString());
  console.log(data);

  if (data.error) {
    if (data.throw) {
      throw new Error("throwing error");
    } else {
      return new Error("returning error");
    }
  }

  return;
};
