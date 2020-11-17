// Package registry provides topics for Device registry operations
package registry

const (
	// actions
	created string = "created"
	updated string = "updated"
	deleted string = "deleted"
	added   string = "added"
	removed string = "removed"

	// thing events
	thing string = "$aws/events/thing/"

	// thing type events
	thingType string = "$aws/events/thingType/"

	// thing type association events
	thingTypeAssociation string = "$aws/events/thingTypeAssociation/thing/"

	// thing group events
	thingGroup string = "$aws/events/thingGroup/"

	// thing group membership events
	thingGroupMembership string = "$aws/events/thingGroupMembership/thingGroup/"

	// thing group hierarchy events
	thingGroupHierarchy string = "$aws/events/thingGroupHierarchy/thingGroup/"
)

//ThingCreated builds topic which listens for thing creation event for specified thingName
func ThingCreated(thingName string) string {
	return thing + thingName + created
}

//ThingUpdated builds topic which listens for thing updating event for specified thingName
func ThingUpdated(thingName string) string {
	return thing + thingName + updated
}

//ThingDeleted builds topic which listens for thing deletion event for specified thingName
func ThingDeleted(thingName string) string {
	return thing + thingName + deleted
}

//ThingTypeCreated builds topic which listens for thing type creation event for specified thingTypeName
func ThingTypeCreated(thingTypeName string) string {
	return thingType + thingTypeName + created
}

//ThingTypeUpdated builds topic which listens for thing type updating event for specified thingTypeName
func ThingTypeUpdated(thingTypeName string) string {
	return thingType + thingTypeName + updated
}

//ThingTypeDeleted builds topic which listens for thing type deletion event for specified thingTypeName
func ThingTypeDeleted(thingTypeName string) string {
	return thingType + thingTypeName + deleted
}

//ThingTypeAssociation builds topic which listens for thing type association event
// for specified thingName & typeName
func ThingTypeAssociation(thingName, typeName string) string {
	return thingTypeAssociation + thingName + "/" + typeName
}

//ThingGroupCreated builds topic which listens for thing group creation event for specified groupName
func ThingGroupCreated(groupName string) string {
	return thingGroup + groupName + created
}

//ThingGroupUpdated builds topic which listens for thing group updating event for specified groupName
func ThingGroupUpdated(groupName string) string {
	return thingGroup + groupName + updated
}

//ThingGroupDeleted builds topic which listens for thing group deleting event for specified groupName
func ThingGroupDeleted(groupName string) string {
	return thingGroup + groupName + deleted
}

//ThingGroupMembershipAdded builds topic which listens for thing group membership addition event
// for specified thingGroupName & thingName
func ThingGroupMembershipAdded(thingGroupName, thingName string) string {
	return thingGroupMembership + thingGroupName + "/thing/" + thingName + added
}

//ThingGroupMembershipRemoved builds topic which listens for thing group membership removal event
// for specified thingGroupName & thingName
func ThingGroupMembershipRemoved(thingGroupName, thingName string) string {
	return thingGroupMembership + thingGroupName + "/thing/" + thingName + removed
}

//ThingGroupHierarchyAdded builds topic which listens for change in thing group hierarchy for addition event
// with specified parentThingGroupName & childThingGroupName
func ThingGroupHierarchyAdded(parentThingGroupName, childThingGroupName string) string {
	return thingGroupHierarchy + parentThingGroupName + "/childThingGroup/" + childThingGroupName + added
}

//ThingGroupHierarchyRemoved builds topic which listens for change in thing group hierarchy for removal event
// with specified parentThingGroupName & childThingGroupName
func ThingGroupHierarchyRemoved(parentThingGroupName, childThingGroupName string) string {
	return thingGroupHierarchy + parentThingGroupName + "/childThingGroup/" + childThingGroupName + removed
}
