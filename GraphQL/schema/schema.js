const graphql = require("graphql");
const { v4: uuid } = require("uuid");

const {
  GraphQLObjectType,
  GraphQLID,
  GraphQLString,
  GraphQLInt,
  GraphQLList,
  GraphQLSchema,
  GraphQLNonNull,
} = graphql;

//Dummy Data for Demo purposes
var stadia = [
  {
    id: "1",
    name: "Chepauk MA Chidambaram",
    capacity: 42000,
    city: "Chennai",
    totalMatchesPlayed: 986,
  },
  {
    id: "2",
    name: "Melbourne Cricket Ground",
    capacity: 30000,
    city: "Melbourne",
    totalMatchesPlayed: 673,
  },
  {
    id: "3",
    name: "The Eden Gardens",
    capacity: 100000,
    city: "Kolkata",
    totalMatchesPlayed: 876,
  },
  {
    id: "4",
    name: "Lord's",
    capacity: 22000,
    city: "London",
    totalMatchesPlayed: 345,
  },
];

var matches = [
  {
    id: "11",
    iccNumber: 1256,
    year: 1978,
    hostTeam: "India",
    guestTeam: "Pakistan",
    venueId: "1",
  },
  {
    id: "12",
    iccNumber: 125,
    year: 1925,
    hostTeam: "England",
    guestTeam: "Australia",
    venueId: "4",
  },
  {
    id: "13",
    iccNumber: 98,
    year: 1925,
    hostTeam: "England",
    guestTeam: "Australia",
    venueId: "2",
  },
  {
    id: "14",
    iccNumber: 225,
    year: 1960,
    hostTeam: "England",
    guestTeam: "India",
    venueId: "3",
  },
  {
    id: "15",
    iccNumber: 2250,
    year: 1994,
    hostTeam: "Australia",
    guestTeam: "India",
    venueId: "2",
  },
  {
    id: "16",
    iccNumber: 1876,
    year: 2012,
    hostTeam: "Australia",
    guestTeam: "India",
    venueId: "3",
  },
];

const StadiumType = new GraphQLObjectType({
  name: "Stadium",
  fields: () => ({
    id: { type: GraphQLID },
    name: { type: GraphQLString },
    capacity: { type: GraphQLInt },
    city: { type: GraphQLString },
    totalMatchesPlayed: { type: GraphQLInt },
    matchesPlayed: {
      type: new GraphQLNonNull(new GraphQLList(MatchType)),
      resolve(parent, args) {
        return matches.filter((m) => m.venueId.localeCompare(parent.id) == 0);
      },
    },
  }),
});

const MatchType = new GraphQLObjectType({
  name: "Match",
  fields: () => ({
    id: { type: GraphQLID },
    iccNumber: { type: GraphQLInt },
    year: { type: GraphQLInt },
    hostTeam: { type: GraphQLString },
    guestTeam: { type: GraphQLString },
    venueId: { type: GraphQLID },
    stadiumInfo: {
      type: StadiumType,
      resolve(parent, args) {
        let stadiaFiltered = stadia.filter(
          (s) => s.id.localeCompare(parent.venueId) == 0
        );
        return stadiaFiltered[0];
      },
    },
  }),
});

const RootQuery = new GraphQLObjectType({
  name: "RootQueryType",
  fields: {
    stadium: {
      type: StadiumType,
      args: { id: { type: new GraphQLNonNull(GraphQLID) } },
      resolve(parent, args) {
        let found = stadia.filter((s) => s.id.localeCompare(args.id) == 0);
        return found[0];
      },
    },
    stadia: {
      type: new GraphQLNonNull(new GraphQLList(StadiumType)),
      resolve(parent, args) {
        return stadia;
      },
    },
    match: {
      type: MatchType,
      args: { id: { type: new GraphQLNonNull(GraphQLID) } },
      resolve(parent, args) {
        let found = matches.filter((m) => m.id.localeCompare(args.id) == 0);
        return found[0];
      },
    },
    matches: {
      type: new GraphQLNonNull(new GraphQLList(MatchType)),
      resolve(parent, args) {
        return matches;
      },
    },
  },
});

const Mutation = new GraphQLObjectType({
  name: "Mutation",
  fields: {
    addStadium: {
      type: StadiumType,
      args: {
        name: { type: new GraphQLNonNull(GraphQLString) },
        capacity: { type: new GraphQLNonNull(GraphQLInt) },
        city: { type: new GraphQLNonNull(GraphQLString) },
        totalMatchesPlayed: { type: new GraphQLNonNull(GraphQLInt) },
      },
      resolve(parent, args) {
        let stadium = {
          id: uuid(),
          name: args.name,
          capacity: args.capacity,
          city: args.city,
          totalMatchesPlayes: args.totalMatchesPlayed,
        };
        stadia.push(stadium);
        return stadium;
      },
    },
    addMatch: {
      type: MatchType,
      args: {
        iccNumber: { type: new GraphQLNonNull(GraphQLInt) },
        hostTeam: { type: new GraphQLNonNull(GraphQLString) },
        guestTeam: { type: new GraphQLNonNull(GraphQLString) },
        totalMatchesPlayed: { type: new GraphQLNonNull(GraphQLInt) },
        year: { type: new GraphQLNonNull(GraphQLInt) },
        venueId: { type: new GraphQLNonNull(GraphQLID) },
      },
      resolve(parent, args) {
        let match = {
          id: uuid(),
          iccNumber: args.iccNumber,
          year: args.year,
          hostTeam: args.hostTeam,
          guestTeam: args.guestTeam,
          venueId: args.venueId,
        };
        console.log(`addMatch: ${JSON.stringify(match)}`);
        matches.push(match);
        return match;
      },
    },
  },
});

module.exports = new GraphQLSchema({
  query: RootQuery,
  mutation: Mutation,
});
