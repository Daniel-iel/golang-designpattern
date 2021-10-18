package main

type NotificationBuilder struct {
	Title    string
	Subtitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) *NotificationBuilder {
	nb.Title = title
	return nb
}

func (nb *NotificationBuilder) SetSubtitle(subtitle string) *NotificationBuilder {
	nb.Subtitle = subtitle
	return nb
}

func (nb *NotificationBuilder) SetMessage(message string) *NotificationBuilder {
	nb.Message = message
	return nb
}

func (nb *NotificationBuilder) SetImage(image string) *NotificationBuilder {
	nb.Image = image
	return nb
}

func (nb *NotificationBuilder) SetIcon(icon string) *NotificationBuilder {
	nb.Icon = icon
	return nb
}

func (nb *NotificationBuilder) SetPriority(pri int) *NotificationBuilder {
	nb.Priority = pri
	return nb
}

func (nb *NotificationBuilder) SetType(notType string) *NotificationBuilder {
	nb.NotType = notType
	return nb
}

func (nb *NotificationBuilder) Build() (*Notification, error) {
	return &Notification{
		title:    nb.Title,
		subtitle: nb.Subtitle,
		message:  nb.Message,
		image:    nb.Image,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}
