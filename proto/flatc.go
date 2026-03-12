package proto

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/mezonai/mezon-protocol/v2/api"
	"github.com/mezonai/mezon-protocol/v2/rtapi"
)

// CreateEnvelope wraps any message offset into an Envelope buffer.
func CreateEnvelope(builder *flatbuffers.Builder, cid string, msgType rtapi.Message, msgOffset flatbuffers.UOffsetT) []byte {
	builder.Reset()

	cidOffset := builder.CreateString(cid)

	rtapi.EnvelopeStart(builder)
	rtapi.EnvelopeAddCid(builder, cidOffset)

	rtapi.EnvelopeAddMessageType(builder, msgType)
	rtapi.EnvelopeAddMessage(builder, msgOffset)

	envelopeOffset := rtapi.EnvelopeEnd(builder)

	builder.Finish(envelopeOffset)
	return builder.FinishedBytes()
}

func GetEnvelopeMessage(env *rtapi.Envelope) interface{} {
	table := new(flatbuffers.Table)
	if !env.Message(table) {
		return nil
	}

	switch env.MessageType() {
	case rtapi.Messagechannel:
		msg := new(rtapi.Channel)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageclan_join:
		msg := new(rtapi.ClanJoin)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagechannel_join:
		msg := new(rtapi.ChannelJoin)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagechannel_leave:
		msg := new(rtapi.ChannelLeave)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagechannel_message:
		msg := new(api.ChannelMessage)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagechannel_message_ack:
		msg := new(rtapi.ChannelMessageAck)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagechannel_message_send:
		msg := new(rtapi.ChannelMessageSend)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagechannel_message_update:
		msg := new(rtapi.ChannelMessageUpdate)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagechannel_message_remove:
		msg := new(rtapi.ChannelMessageRemove)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagechannel_presence_event:
		msg := new(rtapi.ChannelPresenceEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageerror:
		msg := new(rtapi.Error)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagenotifications:
		msg := new(rtapi.Notifications)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagerpc:
		msg := new(api.Rpc)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagestatus:
		msg := new(rtapi.Status)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagestatus_follow:
		msg := new(rtapi.StatusFollow)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagestatus_presence_event:
		msg := new(rtapi.StatusPresenceEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagestatus_unfollow:
		msg := new(rtapi.StatusUnfollow)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagestatus_update:
		msg := new(rtapi.StatusUpdate)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagestream_data:
		msg := new(rtapi.StreamData)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagestream_presence_event:
		msg := new(rtapi.StreamPresenceEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageping:
		msg := new(rtapi.Ping)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagepong:
		msg := new(rtapi.Pong)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagemessage_typing_event:
		msg := new(rtapi.MessageTypingEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagelast_seen_message_event:
		msg := new(rtapi.LastSeenMessageEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagemessage_reaction_event:
		msg := new(api.MessageReaction)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagevoice_joined_event:
		msg := new(rtapi.VoiceJoinedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagevoice_leaved_event:
		msg := new(rtapi.VoiceLeavedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagevoice_started_event:
		msg := new(rtapi.VoiceStartedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagevoice_ended_event:
		msg := new(rtapi.VoiceEndedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagechannel_created_event:
		msg := new(rtapi.ChannelCreatedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagechannel_deleted_event:
		msg := new(rtapi.ChannelDeletedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagechannel_updated_event:
		msg := new(rtapi.ChannelUpdatedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagelast_pin_message_event:
		msg := new(rtapi.LastPinMessageEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagecustom_status_event:
		msg := new(rtapi.CustomStatusEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageuser_channel_added_event:
		msg := new(rtapi.UserChannelAdded)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageuser_channel_removed_event:
		msg := new(rtapi.UserChannelRemoved)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageuser_clan_removed_event:
		msg := new(rtapi.UserClanRemoved)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageclan_updated_event:
		msg := new(rtapi.ClanUpdatedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageclan_profile_updated_event:
		msg := new(rtapi.ClanProfileUpdatedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagecheck_name_existed_event:
		msg := new(rtapi.CheckNameExistedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageuser_profile_updated_event:
		msg := new(rtapi.UserProfileUpdatedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageadd_clan_user_event:
		msg := new(rtapi.AddClanUserEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageclan_event_created:
		msg := new(api.CreateEventRequest)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagerole_assign_event:
		msg := new(rtapi.RoleAssignedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageclan_deleted_event:
		msg := new(rtapi.ClanDeletedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagegive_coffee_event:
		msg := new(api.GiveCoffeeEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagesticker_create_event:
		msg := new(rtapi.StickerCreateEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagesticker_update_event:
		msg := new(rtapi.StickerUpdateEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagesticker_delete_event:
		msg := new(rtapi.StickerDeleteEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagerole_event:
		msg := new(rtapi.RoleEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageevent_emoji:
		msg := new(rtapi.EventEmoji)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagestreaming_joined_event:
		msg := new(rtapi.StreamingJoinedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagestreaming_leaved_event:
		msg := new(rtapi.StreamingLeavedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagestreaming_started_event:
		msg := new(rtapi.StreamingStartedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagestreaming_ended_event:
		msg := new(rtapi.StreamingEndedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagepermission_set_event:
		msg := new(rtapi.PermissionSetEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagepermission_changed_event:
		msg := new(rtapi.PermissionChangedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagetoken_sent_event:
		msg := new(api.TokenSentEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagemessage_button_clicked:
		msg := new(rtapi.MessageButtonClicked)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageunmute_event:
		msg := new(rtapi.UnmuteEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagewebrtc_signaling_fwd:
		msg := new(rtapi.WebrtcSignalingFwd)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagelist_activity:
		msg := new(rtapi.ListActivity)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagedropdown_box_selected:
		msg := new(rtapi.DropdownBoxSelected)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageincoming_call_push:
		msg := new(rtapi.IncomingCallPush)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagesd_topic_event:
		msg := new(rtapi.SdTopicEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagefollow_event:
		msg := new(rtapi.FollowEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagechannel_app_event:
		msg := new(rtapi.ChannelAppEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageuser_status_event:
		msg := new(rtapi.UserStatusEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageremove_friend:
		msg := new(rtapi.RemoveFriend)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagewebhook_event:
		msg := new(api.Webhook)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagenoti_user_channel:
		msg := new(api.NotificationUserChannel)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagejoin_channel_app_data:
		msg := new(rtapi.JoinChannelAppData)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagecanvas_event:
		msg := new(rtapi.ChannelCanvas)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageunpin_message_event:
		msg := new(rtapi.UnpinMessageEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagecategory_event:
		msg := new(rtapi.CategoryEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagehandle_participant_meet_state_event:
		msg := new(rtapi.HandleParticipantMeetStateEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagedelete_account_event:
		msg := new(rtapi.DeleteAccountEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageephemeral_message_send:
		msg := new(rtapi.EphemeralMessageSend)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageblock_friend:
		msg := new(rtapi.BlockFriend)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagevoice_reaction_send:
		msg := new(rtapi.VoiceReactionSend)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagemark_as_read:
		msg := new(rtapi.MarkAsRead)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagelist_data_socket:
		msg := new(rtapi.ListDataSocket)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagequick_menu_event:
		msg := new(rtapi.QuickMenuDataEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageun_block_friend:
		msg := new(rtapi.UnblockFriend)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagemeet_participant_event:
		msg := new(rtapi.MeetParticipantEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messagetransfer_ownership_event:
		msg := new(rtapi.TransferOwnershipEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageadd_friend:
		msg := new(rtapi.AddFriend)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageban_user_event:
		msg := new(rtapi.BannedUserEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageactive_archived_thread:
		msg := new(rtapi.ActiveArchivedThread)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageallow_anonymous_event:
		msg := new(rtapi.AllowAnonymousEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageupdate_localcache_event:
		msg := new(rtapi.UpdateLocalCacheEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageclan_created_event:
		msg := new(rtapi.ClanCreatedEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	case rtapi.Messageaiagent_enabled_event:
		msg := new(rtapi.AIAgentEnabledEvent)
		msg.Init(table.Bytes, table.Pos)
		return msg
	default:
		return nil
	}
}
