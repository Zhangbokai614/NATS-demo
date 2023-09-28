import { connect, StringCodec } from "nats";

// to create a connection to a nats-server:
const nc = await connect({ servers: "192.168.2.251:4223" });

// create a codec
const sc = StringCodec();
// create a simple subscriber and iterate over messages
// matching the subscription
const sub = nc.subscribe("hello");

(async () => {
  for await (const m of sub) {
    console.log(`[${sub.getProcessed()}]: ${sc.decode(m.data)}`);
  }

  console.log("subscription closed");
})();

// nc.publish("hello", sc.encode("world"));
// nc.publish("hello", sc.encode("again"));

// await nc.drain();
