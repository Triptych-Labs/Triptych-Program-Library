import random
import typing
import numpy


class Game:
    hands: numpy.ndarray
    terminated: bool = False
    player_busted: bool = False
    dealer_busted: bool = False

    def __init__(self):
        print("s0")

        cards = [numpy.zeros(5, dtype=int), numpy.zeros(5, dtype=int)]

        for i in range(0, 3):
            turn = int(i / 2)
            rng = random.randint(1, 11)
            cards[turn][i % 2] = rng

        self.hands = cards

    def s2(self):
        print("s2")
        if self.terminated:
            return
        self.terminated = True
        cards = self.hands

        # terminate the game
        #     conduct payout
        #     complete game

        player_hand = numpy.sum(cards[0])
        dealer_hand = numpy.sum(cards[1])

        print()
        print()
        print("Outcome:", self.player_busted, self.dealer_busted)

        if self.player_busted == False and self.dealer_busted:
            print("player wins")
        if self.dealer_busted == False and self.player_busted:
            print("dealer wins")
        if self.dealer_busted == False and self.player_busted == False:
            if player_hand > dealer_hand:
                print("player wins")
            else:
                print("dealer wins")

        print()
        print()

    def sanitize(self):
        print("santizing")
        cards = self.hands

        player_sum = 0
        for card in numpy.where(cards[0] != 0)[0]:
            player_sum += cards[0][card]
        if player_sum > 21:
            self.player_busted = True

        dealer_sum = 0
        for card in numpy.where(cards[1] != 0)[0]:
            dealer_sum += cards[1][card]
        if dealer_sum > 21:
            self.dealer_busted = True

    def s1(
        self,
        enum,
    ):
        if self.terminated:
            raise Exception("terminated")

        print("s1", enum)
        cards = self.hands

        cards_copy = numpy.copy(cards)

        if enum == "hit":
            rng = random.randint(1, 11)
            cards_copy[0][numpy.where(cards_copy[0] == 0)[0][0]] = rng

        if enum == "stand":
            re_entries = 0
            # casino plays soft 17
            while numpy.sum(cards_copy[1]) <= 17:
                if re_entries > 4:
                    break
                rng = random.randint(1, 11)
                cards_copy[1][numpy.where(cards_copy[1] == 0)[0][0]] = rng
                re_entries += 1

        self.hands = cards_copy
        self.sanitize()

        if (
            (self.player_busted == True or self.dealer_busted == True)
            and self.terminated == False
        ) or enum == "stand":
            self.s2()


if __name__ == "__main__":
    print("hello world")

    # [mutable] cards
    game = Game()

    game.s1("hit")
    game.s1("stand")

